// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveCheckInstanceResultWhiteListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RemoveCheckInstanceResultWhiteListResponseBody
	GetRequestId() *string
}

type RemoveCheckInstanceResultWhiteListResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// F156EA41-8EEF-54B2-908B-EAE071XXXXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RemoveCheckInstanceResultWhiteListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RemoveCheckInstanceResultWhiteListResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveCheckInstanceResultWhiteListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RemoveCheckInstanceResultWhiteListResponseBody) SetRequestId(v string) *RemoveCheckInstanceResultWhiteListResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveCheckInstanceResultWhiteListResponseBody) Validate() error {
	return dara.Validate(s)
}
