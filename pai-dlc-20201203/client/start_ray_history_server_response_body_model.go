// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartRayHistoryServerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *StartRayHistoryServerResponseBody
	GetRequestId() *string
}

type StartRayHistoryServerResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StartRayHistoryServerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StartRayHistoryServerResponseBody) GoString() string {
	return s.String()
}

func (s *StartRayHistoryServerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StartRayHistoryServerResponseBody) SetRequestId(v string) *StartRayHistoryServerResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartRayHistoryServerResponseBody) Validate() error {
	return dara.Validate(s)
}
