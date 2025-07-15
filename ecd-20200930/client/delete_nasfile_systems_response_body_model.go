// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteNASFileSystemsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteNASFileSystemsResponseBody
	GetRequestId() *string
}

type DeleteNASFileSystemsResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteNASFileSystemsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteNASFileSystemsResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteNASFileSystemsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteNASFileSystemsResponseBody) SetRequestId(v string) *DeleteNASFileSystemsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteNASFileSystemsResponseBody) Validate() error {
	return dara.Validate(s)
}
