// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDomainResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetDomainResponseBody
	GetRequestId() *string
	SetResult(v map[string]interface{}) *GetDomainResponseBody
	GetResult() map[string]interface{}
}

type GetDomainResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// ABCDEFGH
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The returned result.
	//
	// example:
	//
	// -
	Result map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s GetDomainResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDomainResponseBody) GoString() string {
	return s.String()
}

func (s *GetDomainResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDomainResponseBody) GetResult() map[string]interface{} {
	return s.Result
}

func (s *GetDomainResponseBody) SetRequestId(v string) *GetDomainResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDomainResponseBody) SetResult(v map[string]interface{}) *GetDomainResponseBody {
	s.Result = v
	return s
}

func (s *GetDomainResponseBody) Validate() error {
	return dara.Validate(s)
}
