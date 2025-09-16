// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePublicUrlResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreatePublicUrlResponseBody
	GetRequestId() *string
	SetResult(v map[string]interface{}) *CreatePublicUrlResponseBody
	GetResult() map[string]interface{}
}

type CreatePublicUrlResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 022F36C7-9FB4-5D67-BEBC-3D14B0984463
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The result.
	//
	// example:
	//
	// {}
	Result map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s CreatePublicUrlResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreatePublicUrlResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePublicUrlResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreatePublicUrlResponseBody) GetResult() map[string]interface{} {
	return s.Result
}

func (s *CreatePublicUrlResponseBody) SetRequestId(v string) *CreatePublicUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreatePublicUrlResponseBody) SetResult(v map[string]interface{}) *CreatePublicUrlResponseBody {
	s.Result = v
	return s
}

func (s *CreatePublicUrlResponseBody) Validate() error {
	return dara.Validate(s)
}
