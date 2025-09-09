// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetupTableResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *SetupTableResponseBody
	GetData() *bool
	SetRequestId(v string) *SetupTableResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SetupTableResponseBody
	GetSuccess() *bool
}

type SetupTableResponseBody struct {
	// Specifies whether to use a full table scan.
	//
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// A3140FC7-B78B-4D8E-B0C8-926D28******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SetupTableResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetupTableResponseBody) GoString() string {
	return s.String()
}

func (s *SetupTableResponseBody) GetData() *bool {
	return s.Data
}

func (s *SetupTableResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetupTableResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SetupTableResponseBody) SetData(v bool) *SetupTableResponseBody {
	s.Data = &v
	return s
}

func (s *SetupTableResponseBody) SetRequestId(v string) *SetupTableResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetupTableResponseBody) SetSuccess(v bool) *SetupTableResponseBody {
	s.Success = &v
	return s
}

func (s *SetupTableResponseBody) Validate() error {
	return dara.Validate(s)
}
