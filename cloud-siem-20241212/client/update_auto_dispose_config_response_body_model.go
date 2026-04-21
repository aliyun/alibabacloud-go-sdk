// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAutoDisposeConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateAutoDisposeConfigResponseBody
	GetRequestId() *string
}

type UpdateAutoDisposeConfigResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 6276D891-*****-55B2-87B9-74D413F7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateAutoDisposeConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateAutoDisposeConfigResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAutoDisposeConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateAutoDisposeConfigResponseBody) SetRequestId(v string) *UpdateAutoDisposeConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateAutoDisposeConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
