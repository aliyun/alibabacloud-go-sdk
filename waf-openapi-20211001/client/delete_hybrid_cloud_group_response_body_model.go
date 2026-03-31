// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteHybridCloudGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteHybridCloudGroupResponseBody
	GetRequestId() *string
}

type DeleteHybridCloudGroupResponseBody struct {
	// example:
	//
	// D7861F61-5B61-46CE-A47C-*****60D5EB0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteHybridCloudGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteHybridCloudGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteHybridCloudGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteHybridCloudGroupResponseBody) SetRequestId(v string) *DeleteHybridCloudGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteHybridCloudGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
