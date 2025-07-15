// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDatasetItemResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteDatasetItemResponseBody
	GetRequestId() *string
}

type DeleteDatasetItemResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// FF433E09-663A-5F5D-9DBA-A611********
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDatasetItemResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteDatasetItemResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDatasetItemResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteDatasetItemResponseBody) SetRequestId(v string) *DeleteDatasetItemResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDatasetItemResponseBody) Validate() error {
	return dara.Validate(s)
}
