// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryMscpRiskInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QueryMscpRiskInfoResponseBody
	GetCode() *string
	SetData(v map[string]interface{}) *QueryMscpRiskInfoResponseBody
	GetData() map[string]interface{}
	SetMsg(v string) *QueryMscpRiskInfoResponseBody
	GetMsg() *string
	SetRequestId(v string) *QueryMscpRiskInfoResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryMscpRiskInfoResponseBody
	GetSuccess() *bool
}

type QueryMscpRiskInfoResponseBody struct {
	// Code
	//
	// example:
	//
	// Code
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Data
	//
	// example:
	//
	// Data
	Data map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	// Msg
	//
	// example:
	//
	// Msg
	Msg *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	// Id of the request
	//
	// example:
	//
	// RequestId
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Success
	//
	// example:
	//
	// Success
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryMscpRiskInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryMscpRiskInfoResponseBody) GoString() string {
	return s.String()
}

func (s *QueryMscpRiskInfoResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryMscpRiskInfoResponseBody) GetData() map[string]interface{} {
	return s.Data
}

func (s *QueryMscpRiskInfoResponseBody) GetMsg() *string {
	return s.Msg
}

func (s *QueryMscpRiskInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryMscpRiskInfoResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryMscpRiskInfoResponseBody) SetCode(v string) *QueryMscpRiskInfoResponseBody {
	s.Code = &v
	return s
}

func (s *QueryMscpRiskInfoResponseBody) SetData(v map[string]interface{}) *QueryMscpRiskInfoResponseBody {
	s.Data = v
	return s
}

func (s *QueryMscpRiskInfoResponseBody) SetMsg(v string) *QueryMscpRiskInfoResponseBody {
	s.Msg = &v
	return s
}

func (s *QueryMscpRiskInfoResponseBody) SetRequestId(v string) *QueryMscpRiskInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryMscpRiskInfoResponseBody) SetSuccess(v bool) *QueryMscpRiskInfoResponseBody {
	s.Success = &v
	return s
}

func (s *QueryMscpRiskInfoResponseBody) Validate() error {
	return dara.Validate(s)
}
