package http_server_mappings

import (
	http_server_dtos "github.com/protohedge/protohedge.api/internal/adapters/http_server/dtos"
	"github.com/protohedge/protohedge.api/internal/core/domain"
)

func ToVaultResponseDto(vault *domain.Vault) *http_server_dtos.VaultResponseDto {
	positionManagersDtos := []http_server_dtos.PositionManagerResponseDto{}

	for _, p := range vault.PositionManagers {
		positionManagersDtos = append(positionManagersDtos, ToPositionManagerResponseDto(p))
	}

	return &http_server_dtos.VaultResponseDto{
		VaultAddress:       vault.VaultAddress.String(),
		PositionManagers:   positionManagersDtos,
		VaultWorth:         vault.VaultWorth.String(),
		AvailableLiquidity: vault.AvailableLiquidity.String(),
	}
}

func ToPnlTimestampDto(historicPnl domain.TimePoint) *http_server_dtos.TimePointDto {
	return &http_server_dtos.TimePointDto{
		Timestamp: historicPnl.Timestamp,
		Point:     historicPnl.Point.String(),
	}
}

func ToHistoricPnlDto(historicPnl []domain.TimePoint) *http_server_dtos.HistoricPnlResponseDto {
	historicPnlDto := []*http_server_dtos.TimePointDto{}

	for _, p := range historicPnl {
		historicPnlDto = append(historicPnlDto, ToPnlTimestampDto(p))
	}

	return &http_server_dtos.HistoricPnlResponseDto{
		HistoricPnl: historicPnlDto,
	}
}
