// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateHbaseHaSlbResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateHbaseHaSlbResponseBody
	GetRequestId() *string
}

type CreateHbaseHaSlbResponseBody struct {
	// example:
	//
	// C9D568D9-A59C-4AF2-8FBB-F086A841D58E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateHbaseHaSlbResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateHbaseHaSlbResponseBody) GoString() string {
	return s.String()
}

func (s *CreateHbaseHaSlbResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateHbaseHaSlbResponseBody) SetRequestId(v string) *CreateHbaseHaSlbResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateHbaseHaSlbResponseBody) Validate() error {
	return dara.Validate(s)
}
