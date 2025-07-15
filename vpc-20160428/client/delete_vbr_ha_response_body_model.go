// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteVbrHaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteVbrHaResponseBody
	GetRequestId() *string
}

type DeleteVbrHaResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 4EC47282-1B74-4534-BD0E-403F3EE64CAF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteVbrHaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteVbrHaResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteVbrHaResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteVbrHaResponseBody) SetRequestId(v string) *DeleteVbrHaResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteVbrHaResponseBody) Validate() error {
	return dara.Validate(s)
}
