// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetAppSecretResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ResetAppSecretResponseBody
	GetRequestId() *string
}

type ResetAppSecretResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// CEF72CEB-54B6-4AE8-B225-F876FF7BZ015
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ResetAppSecretResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ResetAppSecretResponseBody) GoString() string {
	return s.String()
}

func (s *ResetAppSecretResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ResetAppSecretResponseBody) SetRequestId(v string) *ResetAppSecretResponseBody {
	s.RequestId = &v
	return s
}

func (s *ResetAppSecretResponseBody) Validate() error {
	return dara.Validate(s)
}
