// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteVulAutoRepairConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteVulAutoRepairConfigResponseBody
	GetRequestId() *string
}

type DeleteVulAutoRepairConfigResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 3956048F-9D73-5EDB-834B-4827BB48****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteVulAutoRepairConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteVulAutoRepairConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteVulAutoRepairConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteVulAutoRepairConfigResponseBody) SetRequestId(v string) *DeleteVulAutoRepairConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteVulAutoRepairConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
