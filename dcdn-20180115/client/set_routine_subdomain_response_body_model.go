// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetRoutineSubdomainResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v map[string]interface{}) *SetRoutineSubdomainResponseBody
	GetContent() map[string]interface{}
	SetRequestId(v string) *SetRoutineSubdomainResponseBody
	GetRequestId() *string
}

type SetRoutineSubdomainResponseBody struct {
	// The message returned, such as ""Status": "OK"".
	//
	// example:
	//
	// "Status": "OK"
	Content map[string]interface{} `json:"Content,omitempty" xml:"Content,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetRoutineSubdomainResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetRoutineSubdomainResponseBody) GoString() string {
	return s.String()
}

func (s *SetRoutineSubdomainResponseBody) GetContent() map[string]interface{} {
	return s.Content
}

func (s *SetRoutineSubdomainResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetRoutineSubdomainResponseBody) SetContent(v map[string]interface{}) *SetRoutineSubdomainResponseBody {
	s.Content = v
	return s
}

func (s *SetRoutineSubdomainResponseBody) SetRequestId(v string) *SetRoutineSubdomainResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetRoutineSubdomainResponseBody) Validate() error {
	return dara.Validate(s)
}
