// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCutOverReplicationJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CutOverReplicationJobResponseBody
	GetRequestId() *string
}

type CutOverReplicationJobResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CutOverReplicationJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CutOverReplicationJobResponseBody) GoString() string {
	return s.String()
}

func (s *CutOverReplicationJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CutOverReplicationJobResponseBody) SetRequestId(v string) *CutOverReplicationJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *CutOverReplicationJobResponseBody) Validate() error {
	return dara.Validate(s)
}
