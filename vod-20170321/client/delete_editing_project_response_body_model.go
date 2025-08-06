// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEditingProjectResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteEditingProjectResponseBody
	GetRequestId() *string
}

type DeleteEditingProjectResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 25818875-5F78-4A*****F6-D7393642CA58
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteEditingProjectResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteEditingProjectResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteEditingProjectResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteEditingProjectResponseBody) SetRequestId(v string) *DeleteEditingProjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteEditingProjectResponseBody) Validate() error {
	return dara.Validate(s)
}
