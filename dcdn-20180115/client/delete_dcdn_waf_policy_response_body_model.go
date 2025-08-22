// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDcdnWafPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteDcdnWafPolicyResponseBody
	GetRequestId() *string
}

type DeleteDcdnWafPolicyResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// CB1A380B-09F0-41BB-B084-72F8FD6DA2FE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDcdnWafPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteDcdnWafPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDcdnWafPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteDcdnWafPolicyResponseBody) SetRequestId(v string) *DeleteDcdnWafPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDcdnWafPolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
