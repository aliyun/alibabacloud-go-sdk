// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRadioRunHistoryTimeSerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApMac(v string) *GetRadioRunHistoryTimeSerRequest
	GetApMac() *string
	SetAppCode(v string) *GetRadioRunHistoryTimeSerRequest
	GetAppCode() *string
	SetAppName(v string) *GetRadioRunHistoryTimeSerRequest
	GetAppName() *string
	SetEnd(v int64) *GetRadioRunHistoryTimeSerRequest
	GetEnd() *int64
	SetStart(v int64) *GetRadioRunHistoryTimeSerRequest
	GetStart() *int64
}

type GetRadioRunHistoryTimeSerRequest struct {
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

func (s GetRadioRunHistoryTimeSerRequest) String() string {
	return dara.Prettify(s)
}

func (s GetRadioRunHistoryTimeSerRequest) GoString() string {
	return s.String()
}

func (s *GetRadioRunHistoryTimeSerRequest) GetApMac() *string {
	return s.ApMac
}

func (s *GetRadioRunHistoryTimeSerRequest) GetAppCode() *string {
	return s.AppCode
}

func (s *GetRadioRunHistoryTimeSerRequest) GetAppName() *string {
	return s.AppName
}

func (s *GetRadioRunHistoryTimeSerRequest) GetEnd() *int64 {
	return s.End
}

func (s *GetRadioRunHistoryTimeSerRequest) GetStart() *int64 {
	return s.Start
}

func (s *GetRadioRunHistoryTimeSerRequest) SetApMac(v string) *GetRadioRunHistoryTimeSerRequest {
	s.ApMac = &v
	return s
}

func (s *GetRadioRunHistoryTimeSerRequest) SetAppCode(v string) *GetRadioRunHistoryTimeSerRequest {
	s.AppCode = &v
	return s
}

func (s *GetRadioRunHistoryTimeSerRequest) SetAppName(v string) *GetRadioRunHistoryTimeSerRequest {
	s.AppName = &v
	return s
}

func (s *GetRadioRunHistoryTimeSerRequest) SetEnd(v int64) *GetRadioRunHistoryTimeSerRequest {
	s.End = &v
	return s
}

func (s *GetRadioRunHistoryTimeSerRequest) SetStart(v int64) *GetRadioRunHistoryTimeSerRequest {
	s.Start = &v
	return s
}

func (s *GetRadioRunHistoryTimeSerRequest) Validate() error {
	return dara.Validate(s)
}
