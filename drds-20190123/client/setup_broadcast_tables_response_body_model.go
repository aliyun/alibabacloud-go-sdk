// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetupBroadcastTablesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *SetupBroadcastTablesResponseBody
	GetData() *bool
	SetRequestId(v string) *SetupBroadcastTablesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SetupBroadcastTablesResponseBody
	GetSuccess() *bool
}

type SetupBroadcastTablesResponseBody struct {
	// Indicates whether the broadcast table is configured.
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
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SetupBroadcastTablesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetupBroadcastTablesResponseBody) GoString() string {
	return s.String()
}

func (s *SetupBroadcastTablesResponseBody) GetData() *bool {
	return s.Data
}

func (s *SetupBroadcastTablesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetupBroadcastTablesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SetupBroadcastTablesResponseBody) SetData(v bool) *SetupBroadcastTablesResponseBody {
	s.Data = &v
	return s
}

func (s *SetupBroadcastTablesResponseBody) SetRequestId(v string) *SetupBroadcastTablesResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetupBroadcastTablesResponseBody) SetSuccess(v bool) *SetupBroadcastTablesResponseBody {
	s.Success = &v
	return s
}

func (s *SetupBroadcastTablesResponseBody) Validate() error {
	return dara.Validate(s)
}
