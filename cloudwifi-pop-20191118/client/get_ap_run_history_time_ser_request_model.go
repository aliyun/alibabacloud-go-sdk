// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetApRunHistoryTimeSerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApMac(v string) *GetApRunHistoryTimeSerRequest
	GetApMac() *string
	SetAppCode(v string) *GetApRunHistoryTimeSerRequest
	GetAppCode() *string
	SetAppName(v string) *GetApRunHistoryTimeSerRequest
	GetAppName() *string
	SetEnd(v int64) *GetApRunHistoryTimeSerRequest
	GetEnd() *int64
	SetStart(v int64) *GetApRunHistoryTimeSerRequest
	GetStart() *int64
}

type GetApRunHistoryTimeSerRequest struct {
	// This parameter is required.
	ApMac *string `json:"ApMac,omitempty" xml:"ApMac,omitempty"`
	// This parameter is required.
	AppCode *string `json:"AppCode,omitempty" xml:"AppCode,omitempty"`
	// This parameter is required.
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// This parameter is required.
	End *int64 `json:"End,omitempty" xml:"End,omitempty"`
	// This parameter is required.
	Start *int64 `json:"Start,omitempty" xml:"Start,omitempty"`
}

func (s GetApRunHistoryTimeSerRequest) String() string {
	return dara.Prettify(s)
}

func (s GetApRunHistoryTimeSerRequest) GoString() string {
	return s.String()
}

func (s *GetApRunHistoryTimeSerRequest) GetApMac() *string {
	return s.ApMac
}

func (s *GetApRunHistoryTimeSerRequest) GetAppCode() *string {
	return s.AppCode
}

func (s *GetApRunHistoryTimeSerRequest) GetAppName() *string {
	return s.AppName
}

func (s *GetApRunHistoryTimeSerRequest) GetEnd() *int64 {
	return s.End
}

func (s *GetApRunHistoryTimeSerRequest) GetStart() *int64 {
	return s.Start
}

func (s *GetApRunHistoryTimeSerRequest) SetApMac(v string) *GetApRunHistoryTimeSerRequest {
	s.ApMac = &v
	return s
}

func (s *GetApRunHistoryTimeSerRequest) SetAppCode(v string) *GetApRunHistoryTimeSerRequest {
	s.AppCode = &v
	return s
}

func (s *GetApRunHistoryTimeSerRequest) SetAppName(v string) *GetApRunHistoryTimeSerRequest {
	s.AppName = &v
	return s
}

func (s *GetApRunHistoryTimeSerRequest) SetEnd(v int64) *GetApRunHistoryTimeSerRequest {
	s.End = &v
	return s
}

func (s *GetApRunHistoryTimeSerRequest) SetStart(v int64) *GetApRunHistoryTimeSerRequest {
	s.Start = &v
	return s
}

func (s *GetApRunHistoryTimeSerRequest) Validate() error {
	return dara.Validate(s)
}
