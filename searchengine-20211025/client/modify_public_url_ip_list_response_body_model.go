// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyPublicUrlIpListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyPublicUrlIpListResponseBody
	GetRequestId() *string
	SetResult(v map[string]interface{}) *ModifyPublicUrlIpListResponseBody
	GetResult() map[string]interface{}
}

type ModifyPublicUrlIpListResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// E45380E8-994A-5402-9806-F114B3295FCF
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The result.
	//
	// example:
	//
	// {}
	Result map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s ModifyPublicUrlIpListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyPublicUrlIpListResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyPublicUrlIpListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyPublicUrlIpListResponseBody) GetResult() map[string]interface{} {
	return s.Result
}

func (s *ModifyPublicUrlIpListResponseBody) SetRequestId(v string) *ModifyPublicUrlIpListResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyPublicUrlIpListResponseBody) SetResult(v map[string]interface{}) *ModifyPublicUrlIpListResponseBody {
	s.Result = v
	return s
}

func (s *ModifyPublicUrlIpListResponseBody) Validate() error {
	return dara.Validate(s)
}
