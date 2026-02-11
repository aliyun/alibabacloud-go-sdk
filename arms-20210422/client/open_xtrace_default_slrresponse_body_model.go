// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOpenXtraceDefaultSLRResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v string) *OpenXtraceDefaultSLRResponseBody
	GetData() *string
	SetRequestId(v string) *OpenXtraceDefaultSLRResponseBody
	GetRequestId() *string
}

type OpenXtraceDefaultSLRResponseBody struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s OpenXtraceDefaultSLRResponseBody) String() string {
	return dara.Prettify(s)
}

func (s OpenXtraceDefaultSLRResponseBody) GoString() string {
	return s.String()
}

func (s *OpenXtraceDefaultSLRResponseBody) GetData() *string {
	return s.Data
}

func (s *OpenXtraceDefaultSLRResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *OpenXtraceDefaultSLRResponseBody) SetData(v string) *OpenXtraceDefaultSLRResponseBody {
	s.Data = &v
	return s
}

func (s *OpenXtraceDefaultSLRResponseBody) SetRequestId(v string) *OpenXtraceDefaultSLRResponseBody {
	s.RequestId = &v
	return s
}

func (s *OpenXtraceDefaultSLRResponseBody) Validate() error {
	return dara.Validate(s)
}
