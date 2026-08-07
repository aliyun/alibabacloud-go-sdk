// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAppAgentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *DeleteAppAgentResponseBody
	GetData() *bool
	SetRequestId(v string) *DeleteAppAgentResponseBody
	GetRequestId() *string
}

type DeleteAppAgentResponseBody struct {
	// The returned data.
	//
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// The ID assigned by the backend to uniquely identify a request. It can be used to troubleshoot issues.
	//
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteAppAgentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteAppAgentResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAppAgentResponseBody) GetData() *bool {
	return s.Data
}

func (s *DeleteAppAgentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteAppAgentResponseBody) SetData(v bool) *DeleteAppAgentResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteAppAgentResponseBody) SetRequestId(v string) *DeleteAppAgentResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAppAgentResponseBody) Validate() error {
	return dara.Validate(s)
}
