// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPurgeInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *PurgeInstanceResponseBody
	GetRequestId() *string
}

type PurgeInstanceResponseBody struct {
	// example:
	//
	// 276F899F-E952-496F-81B8-BD46D86854E3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s PurgeInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PurgeInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *PurgeInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PurgeInstanceResponseBody) SetRequestId(v string) *PurgeInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *PurgeInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
