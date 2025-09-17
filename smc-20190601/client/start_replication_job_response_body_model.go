// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartReplicationJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *StartReplicationJobResponseBody
	GetRequestId() *string
}

type StartReplicationJobResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StartReplicationJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StartReplicationJobResponseBody) GoString() string {
	return s.String()
}

func (s *StartReplicationJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StartReplicationJobResponseBody) SetRequestId(v string) *StartReplicationJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartReplicationJobResponseBody) Validate() error {
	return dara.Validate(s)
}
