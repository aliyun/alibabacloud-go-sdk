// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindVbrResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *BindVbrResponseBody
	GetRequestId() *string
}

type BindVbrResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 09A2010F-602B-4EC6-A60F-9914AAE2DCA0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s BindVbrResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BindVbrResponseBody) GoString() string {
	return s.String()
}

func (s *BindVbrResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BindVbrResponseBody) SetRequestId(v string) *BindVbrResponseBody {
	s.RequestId = &v
	return s
}

func (s *BindVbrResponseBody) Validate() error {
	return dara.Validate(s)
}
