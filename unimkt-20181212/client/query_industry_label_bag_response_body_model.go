// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryIndustryLabelBagResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *QueryIndustryLabelBagResponseBody
	GetCode() *int32
	SetData(v []map[string]interface{}) *QueryIndustryLabelBagResponseBody
	GetData() []map[string]interface{}
	SetErrorMsg(v string) *QueryIndustryLabelBagResponseBody
	GetErrorMsg() *string
	SetRequestId(v string) *QueryIndustryLabelBagResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryIndustryLabelBagResponseBody
	GetSuccess() *bool
}

type QueryIndustryLabelBagResponseBody struct {
	Code      *int32                   `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      []map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	ErrorMsg  *string                  `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	RequestId *string                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                    `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryIndustryLabelBagResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryIndustryLabelBagResponseBody) GoString() string {
	return s.String()
}

func (s *QueryIndustryLabelBagResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *QueryIndustryLabelBagResponseBody) GetData() []map[string]interface{} {
	return s.Data
}

func (s *QueryIndustryLabelBagResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *QueryIndustryLabelBagResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryIndustryLabelBagResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryIndustryLabelBagResponseBody) SetCode(v int32) *QueryIndustryLabelBagResponseBody {
	s.Code = &v
	return s
}

func (s *QueryIndustryLabelBagResponseBody) SetData(v []map[string]interface{}) *QueryIndustryLabelBagResponseBody {
	s.Data = v
	return s
}

func (s *QueryIndustryLabelBagResponseBody) SetErrorMsg(v string) *QueryIndustryLabelBagResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *QueryIndustryLabelBagResponseBody) SetRequestId(v string) *QueryIndustryLabelBagResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryIndustryLabelBagResponseBody) SetSuccess(v bool) *QueryIndustryLabelBagResponseBody {
	s.Success = &v
	return s
}

func (s *QueryIndustryLabelBagResponseBody) Validate() error {
	return dara.Validate(s)
}
