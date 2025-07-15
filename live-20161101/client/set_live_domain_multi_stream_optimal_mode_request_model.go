// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetLiveDomainMultiStreamOptimalModeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *SetLiveDomainMultiStreamOptimalModeRequest
	GetAppName() *string
	SetDomain(v string) *SetLiveDomainMultiStreamOptimalModeRequest
	GetDomain() *string
	SetOptimalMode(v string) *SetLiveDomainMultiStreamOptimalModeRequest
	GetOptimalMode() *string
	SetOwnerId(v int64) *SetLiveDomainMultiStreamOptimalModeRequest
	GetOwnerId() *int64
	SetStreamName(v string) *SetLiveDomainMultiStreamOptimalModeRequest
	GetStreamName() *string
}

type SetLiveDomainMultiStreamOptimalModeRequest struct {
	// The application name.
	//
	// This parameter is required.
	//
	// example:
	//
	// testapp
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The main streaming domain.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// Specifies whether to enable the auto mode of dual-stream disaster recovery. Valid values:
	//
	// 	- **on**: enables the auto mode.
	//
	// 	- **off**: disables the auto mode.
	//
	// This parameter is required.
	//
	// example:
	//
	// on
	OptimalMode *string `json:"OptimalMode,omitempty" xml:"OptimalMode,omitempty"`
	OwnerId     *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The name of the live stream.
	//
	// This parameter is required.
	//
	// example:
	//
	// teststream
	StreamName *string `json:"StreamName,omitempty" xml:"StreamName,omitempty"`
}

func (s SetLiveDomainMultiStreamOptimalModeRequest) String() string {
	return dara.Prettify(s)
}

func (s SetLiveDomainMultiStreamOptimalModeRequest) GoString() string {
	return s.String()
}

func (s *SetLiveDomainMultiStreamOptimalModeRequest) GetAppName() *string {
	return s.AppName
}

func (s *SetLiveDomainMultiStreamOptimalModeRequest) GetDomain() *string {
	return s.Domain
}

func (s *SetLiveDomainMultiStreamOptimalModeRequest) GetOptimalMode() *string {
	return s.OptimalMode
}

func (s *SetLiveDomainMultiStreamOptimalModeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *SetLiveDomainMultiStreamOptimalModeRequest) GetStreamName() *string {
	return s.StreamName
}

func (s *SetLiveDomainMultiStreamOptimalModeRequest) SetAppName(v string) *SetLiveDomainMultiStreamOptimalModeRequest {
	s.AppName = &v
	return s
}

func (s *SetLiveDomainMultiStreamOptimalModeRequest) SetDomain(v string) *SetLiveDomainMultiStreamOptimalModeRequest {
	s.Domain = &v
	return s
}

func (s *SetLiveDomainMultiStreamOptimalModeRequest) SetOptimalMode(v string) *SetLiveDomainMultiStreamOptimalModeRequest {
	s.OptimalMode = &v
	return s
}

func (s *SetLiveDomainMultiStreamOptimalModeRequest) SetOwnerId(v int64) *SetLiveDomainMultiStreamOptimalModeRequest {
	s.OwnerId = &v
	return s
}

func (s *SetLiveDomainMultiStreamOptimalModeRequest) SetStreamName(v string) *SetLiveDomainMultiStreamOptimalModeRequest {
	s.StreamName = &v
	return s
}

func (s *SetLiveDomainMultiStreamOptimalModeRequest) Validate() error {
	return dara.Validate(s)
}
