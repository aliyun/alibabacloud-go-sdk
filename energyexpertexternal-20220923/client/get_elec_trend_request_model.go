// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetElecTrendRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetElecTrendRequest
	GetCode() *string
	SetYearList(v []*int32) *GetElecTrendRequest
	GetYearList() []*int32
}

type GetElecTrendRequest struct {
	// The enterprise code.
	//
	// This parameter is required.
	//
	// example:
	//
	// C-20240115-3
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// List of years.
	//
	// This parameter is required.
	YearList []*int32 `json:"yearList,omitempty" xml:"yearList,omitempty" type:"Repeated"`
}

func (s GetElecTrendRequest) String() string {
	return dara.Prettify(s)
}

func (s GetElecTrendRequest) GoString() string {
	return s.String()
}

func (s *GetElecTrendRequest) GetCode() *string {
	return s.Code
}

func (s *GetElecTrendRequest) GetYearList() []*int32 {
	return s.YearList
}

func (s *GetElecTrendRequest) SetCode(v string) *GetElecTrendRequest {
	s.Code = &v
	return s
}

func (s *GetElecTrendRequest) SetYearList(v []*int32) *GetElecTrendRequest {
	s.YearList = v
	return s
}

func (s *GetElecTrendRequest) Validate() error {
	return dara.Validate(s)
}
