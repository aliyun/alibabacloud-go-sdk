// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveSignatureApisResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RemoveSignatureApisResponseBody
	GetRequestId() *string
}

type RemoveSignatureApisResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// CEF72CEB-54B6-4AE8-B225-F876FF7BZ004
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RemoveSignatureApisResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RemoveSignatureApisResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveSignatureApisResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RemoveSignatureApisResponseBody) SetRequestId(v string) *RemoveSignatureApisResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveSignatureApisResponseBody) Validate() error {
	return dara.Validate(s)
}
