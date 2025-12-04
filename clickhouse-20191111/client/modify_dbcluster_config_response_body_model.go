// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBClusterConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyDBClusterConfigResponseBody
	GetRequestId() *string
}

type ModifyDBClusterConfigResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// A9AA1E0A-2AEE-5847-87CF-E4FDC0E66052
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDBClusterConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBClusterConfigResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyDBClusterConfigResponseBody) SetRequestId(v string) *ModifyDBClusterConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDBClusterConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
