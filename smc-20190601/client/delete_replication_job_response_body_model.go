// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteReplicationJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteReplicationJobResponseBody
	GetRequestId() *string
}

type DeleteReplicationJobResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteReplicationJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteReplicationJobResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteReplicationJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteReplicationJobResponseBody) SetRequestId(v string) *DeleteReplicationJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteReplicationJobResponseBody) Validate() error {
	return dara.Validate(s)
}
