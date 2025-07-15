// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteApiResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteApiResponseBody
	GetRequestId() *string
}

type DeleteApiResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// CEF72CEB-54B6-4AE8-B225-F876FF7BZ017
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteApiResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteApiResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteApiResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteApiResponseBody) SetRequestId(v string) *DeleteApiResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteApiResponseBody) Validate() error {
	return dara.Validate(s)
}
