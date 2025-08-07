// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBClusterTDEResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyDBClusterTDEResponseBody
	GetRequestId() *string
}

type ModifyDBClusterTDEResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 5F859238-2A36-4A8D-BD0F-732112******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDBClusterTDEResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBClusterTDEResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterTDEResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyDBClusterTDEResponseBody) SetRequestId(v string) *ModifyDBClusterTDEResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDBClusterTDEResponseBody) Validate() error {
	return dara.Validate(s)
}
