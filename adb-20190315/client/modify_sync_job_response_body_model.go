// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySyncJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifySyncJobResponseBody
	GetRequestId() *string
}

type ModifySyncJobResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 5AD3D9DF-614F-5B97-9522-A2406A432012
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifySyncJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifySyncJobResponseBody) GoString() string {
	return s.String()
}

func (s *ModifySyncJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifySyncJobResponseBody) SetRequestId(v string) *ModifySyncJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifySyncJobResponseBody) Validate() error {
	return dara.Validate(s)
}
