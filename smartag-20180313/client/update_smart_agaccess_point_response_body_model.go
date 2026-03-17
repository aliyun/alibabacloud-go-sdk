// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSmartAGAccessPointResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateSmartAGAccessPointResponseBody
	GetRequestId() *string
}

type UpdateSmartAGAccessPointResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// E26DBAAE-A796-4A48-98B4-B45AFCD1F299
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateSmartAGAccessPointResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateSmartAGAccessPointResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateSmartAGAccessPointResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateSmartAGAccessPointResponseBody) SetRequestId(v string) *UpdateSmartAGAccessPointResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateSmartAGAccessPointResponseBody) Validate() error {
	return dara.Validate(s)
}
