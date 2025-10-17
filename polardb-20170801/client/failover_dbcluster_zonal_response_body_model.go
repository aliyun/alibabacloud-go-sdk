// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFailoverDBClusterZonalResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *FailoverDBClusterZonalResponseBody
	GetRequestId() *string
}

type FailoverDBClusterZonalResponseBody struct {
	// example:
	//
	// 925B84D9-CA72-432C-95CF-738C22******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s FailoverDBClusterZonalResponseBody) String() string {
	return dara.Prettify(s)
}

func (s FailoverDBClusterZonalResponseBody) GoString() string {
	return s.String()
}

func (s *FailoverDBClusterZonalResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *FailoverDBClusterZonalResponseBody) SetRequestId(v string) *FailoverDBClusterZonalResponseBody {
	s.RequestId = &v
	return s
}

func (s *FailoverDBClusterZonalResponseBody) Validate() error {
	return dara.Validate(s)
}
