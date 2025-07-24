// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunApiTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v string) *RunApiTemplateResponseBody
	GetData() *string
	SetRequestId(v string) *RunApiTemplateResponseBody
	GetRequestId() *string
}

type RunApiTemplateResponseBody struct {
	// example:
	//
	// {\\"clusterId\\":\\"c-b7be171f1928****\\",\\"operationId\\":\\"op-61126efe629d****\\"}
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// 请求ID。
	//
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C8944****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RunApiTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RunApiTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *RunApiTemplateResponseBody) GetData() *string {
	return s.Data
}

func (s *RunApiTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RunApiTemplateResponseBody) SetData(v string) *RunApiTemplateResponseBody {
	s.Data = &v
	return s
}

func (s *RunApiTemplateResponseBody) SetRequestId(v string) *RunApiTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *RunApiTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}
