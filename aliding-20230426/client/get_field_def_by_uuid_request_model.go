// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFieldDefByUuidRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppType(v string) *GetFieldDefByUuidRequest
	GetAppType() *string
	SetFormUuid(v string) *GetFieldDefByUuidRequest
	GetFormUuid() *string
	SetSystemToken(v string) *GetFieldDefByUuidRequest
	GetSystemToken() *string
}

type GetFieldDefByUuidRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// APP_PBKT0xxx
	AppType *string `json:"AppType,omitempty" xml:"AppType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// FORM-xxxxx
	FormUuid *string `json:"FormUuid,omitempty" xml:"FormUuid,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// hexxxx
	SystemToken *string `json:"SystemToken,omitempty" xml:"SystemToken,omitempty"`
}

func (s GetFieldDefByUuidRequest) String() string {
	return dara.Prettify(s)
}

func (s GetFieldDefByUuidRequest) GoString() string {
	return s.String()
}

func (s *GetFieldDefByUuidRequest) GetAppType() *string {
	return s.AppType
}

func (s *GetFieldDefByUuidRequest) GetFormUuid() *string {
	return s.FormUuid
}

func (s *GetFieldDefByUuidRequest) GetSystemToken() *string {
	return s.SystemToken
}

func (s *GetFieldDefByUuidRequest) SetAppType(v string) *GetFieldDefByUuidRequest {
	s.AppType = &v
	return s
}

func (s *GetFieldDefByUuidRequest) SetFormUuid(v string) *GetFieldDefByUuidRequest {
	s.FormUuid = &v
	return s
}

func (s *GetFieldDefByUuidRequest) SetSystemToken(v string) *GetFieldDefByUuidRequest {
	s.SystemToken = &v
	return s
}

func (s *GetFieldDefByUuidRequest) Validate() error {
	return dara.Validate(s)
}
