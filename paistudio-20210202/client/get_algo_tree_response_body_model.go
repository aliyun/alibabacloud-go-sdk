// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAlgoTreeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v map[string]interface{}) *GetAlgoTreeResponseBody
	GetData() map[string]interface{}
	SetRequestId(v string) *GetAlgoTreeResponseBody
	GetRequestId() *string
}

type GetAlgoTreeResponseBody struct {
	// example:
	//
	// [{}]
	Data map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 8CAA642F-EFEB-5F87-8F2F-ACD58B15FA03
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetAlgoTreeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAlgoTreeResponseBody) GoString() string {
	return s.String()
}

func (s *GetAlgoTreeResponseBody) GetData() map[string]interface{} {
	return s.Data
}

func (s *GetAlgoTreeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAlgoTreeResponseBody) SetData(v map[string]interface{}) *GetAlgoTreeResponseBody {
	s.Data = v
	return s
}

func (s *GetAlgoTreeResponseBody) SetRequestId(v string) *GetAlgoTreeResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAlgoTreeResponseBody) Validate() error {
	return dara.Validate(s)
}
