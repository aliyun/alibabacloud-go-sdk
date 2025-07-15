// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVbrHaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateVbrHaResponseBody
	GetRequestId() *string
	SetVbrHaId(v string) *CreateVbrHaResponseBody
	GetVbrHaId() *string
}

type CreateVbrHaResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 4EC47282-1B74-4534-BD0E-403F3EE64CAF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the VBR failover group.
	//
	// example:
	//
	// vbrha-sa1sxheuxtd98****
	VbrHaId *string `json:"VbrHaId,omitempty" xml:"VbrHaId,omitempty"`
}

func (s CreateVbrHaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateVbrHaResponseBody) GoString() string {
	return s.String()
}

func (s *CreateVbrHaResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateVbrHaResponseBody) GetVbrHaId() *string {
	return s.VbrHaId
}

func (s *CreateVbrHaResponseBody) SetRequestId(v string) *CreateVbrHaResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateVbrHaResponseBody) SetVbrHaId(v string) *CreateVbrHaResponseBody {
	s.VbrHaId = &v
	return s
}

func (s *CreateVbrHaResponseBody) Validate() error {
	return dara.Validate(s)
}
