// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSavedQueryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteSavedQueryResponseBody
	GetRequestId() *string
}

type DeleteSavedQueryResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// D170D58E-6256-5344-8F5E-922EC9ECB7EA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteSavedQueryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteSavedQueryResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSavedQueryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteSavedQueryResponseBody) SetRequestId(v string) *DeleteSavedQueryResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteSavedQueryResponseBody) Validate() error {
	return dara.Validate(s)
}
