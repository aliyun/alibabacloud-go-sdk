// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddVsPullStreamInfoConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AddVsPullStreamInfoConfigResponseBody
	GetRequestId() *string
}

type AddVsPullStreamInfoConfigResponseBody struct {
	// example:
	//
	// 3CB843A9-DD34-4881-B8D6-B0D539D111E4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddVsPullStreamInfoConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddVsPullStreamInfoConfigResponseBody) GoString() string {
	return s.String()
}

func (s *AddVsPullStreamInfoConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddVsPullStreamInfoConfigResponseBody) SetRequestId(v string) *AddVsPullStreamInfoConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddVsPullStreamInfoConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
