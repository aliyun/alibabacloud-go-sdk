// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListParamsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetParams(v []*ListParamsResponseBodyParams) *ListParamsResponseBody
	GetParams() []*ListParamsResponseBodyParams
	SetRequestId(v string) *ListParamsResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListParamsResponseBody
	GetTotalCount() *int64
}

type ListParamsResponseBody struct {
	Params []*ListParamsResponseBodyParams `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	// Id of the request
	//
	// example:
	//
	// A2D07551-38DA-531E-9B22-877D1D86A579
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 10
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListParamsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListParamsResponseBody) GoString() string {
	return s.String()
}

func (s *ListParamsResponseBody) GetParams() []*ListParamsResponseBodyParams {
	return s.Params
}

func (s *ListParamsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListParamsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListParamsResponseBody) SetParams(v []*ListParamsResponseBodyParams) *ListParamsResponseBody {
	s.Params = v
	return s
}

func (s *ListParamsResponseBody) SetRequestId(v string) *ListParamsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListParamsResponseBody) SetTotalCount(v int64) *ListParamsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListParamsResponseBody) Validate() error {
	if s.Params != nil {
		for _, item := range s.Params {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListParamsResponseBodyParams struct {
	// example:
	//
	// Daily
	Environment *string `json:"Environment,omitempty" xml:"Environment,omitempty"`
	// example:
	//
	// 2021-12-15T23:24:33.132+08:00
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	// example:
	//
	// home
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 4
	ParamId *string `json:"ParamId,omitempty" xml:"ParamId,omitempty"`
	// example:
	//
	// house
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListParamsResponseBodyParams) String() string {
	return dara.Prettify(s)
}

func (s ListParamsResponseBodyParams) GoString() string {
	return s.String()
}

func (s *ListParamsResponseBodyParams) GetEnvironment() *string {
	return s.Environment
}

func (s *ListParamsResponseBodyParams) GetGmtModifiedTime() *string {
	return s.GmtModifiedTime
}

func (s *ListParamsResponseBodyParams) GetName() *string {
	return s.Name
}

func (s *ListParamsResponseBodyParams) GetParamId() *string {
	return s.ParamId
}

func (s *ListParamsResponseBodyParams) GetValue() *string {
	return s.Value
}

func (s *ListParamsResponseBodyParams) SetEnvironment(v string) *ListParamsResponseBodyParams {
	s.Environment = &v
	return s
}

func (s *ListParamsResponseBodyParams) SetGmtModifiedTime(v string) *ListParamsResponseBodyParams {
	s.GmtModifiedTime = &v
	return s
}

func (s *ListParamsResponseBodyParams) SetName(v string) *ListParamsResponseBodyParams {
	s.Name = &v
	return s
}

func (s *ListParamsResponseBodyParams) SetParamId(v string) *ListParamsResponseBodyParams {
	s.ParamId = &v
	return s
}

func (s *ListParamsResponseBodyParams) SetValue(v string) *ListParamsResponseBodyParams {
	s.Value = &v
	return s
}

func (s *ListParamsResponseBodyParams) Validate() error {
	return dara.Validate(s)
}
