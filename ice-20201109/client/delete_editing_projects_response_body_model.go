// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEditingProjectsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteEditingProjectsResponseBody
	GetRequestId() *string
}

type DeleteEditingProjectsResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// ****25818875-5F78-4AF6-D7393642CA58****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteEditingProjectsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteEditingProjectsResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteEditingProjectsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteEditingProjectsResponseBody) SetRequestId(v string) *DeleteEditingProjectsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteEditingProjectsResponseBody) Validate() error {
	return dara.Validate(s)
}
