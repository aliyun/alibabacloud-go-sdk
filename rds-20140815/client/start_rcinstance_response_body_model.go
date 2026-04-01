// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartRCInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *StartRCInstanceResponseBody
	GetRequestId() *string
}

type StartRCInstanceResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 8B993DA9-5272-5414-94E3-4CA8BA0146C2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StartRCInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StartRCInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *StartRCInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StartRCInstanceResponseBody) SetRequestId(v string) *StartRCInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartRCInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
