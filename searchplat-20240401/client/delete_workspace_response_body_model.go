// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteWorkspaceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteWorkspaceResponseBody
	GetRequestId() *string
	SetResult(v *DeleteWorkspaceResponseBodyResult) *DeleteWorkspaceResponseBody
	GetResult() *DeleteWorkspaceResponseBodyResult
}

type DeleteWorkspaceResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 5950143C-B8F0-5758-A08A-66F302FD587F
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The returned result.
	Result *DeleteWorkspaceResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s DeleteWorkspaceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteWorkspaceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteWorkspaceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteWorkspaceResponseBody) GetResult() *DeleteWorkspaceResponseBodyResult {
	return s.Result
}

func (s *DeleteWorkspaceResponseBody) SetRequestId(v string) *DeleteWorkspaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteWorkspaceResponseBody) SetResult(v *DeleteWorkspaceResponseBodyResult) *DeleteWorkspaceResponseBody {
	s.Result = v
	return s
}

func (s *DeleteWorkspaceResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeleteWorkspaceResponseBodyResult struct {
	// The instance ID.
	//
	// example:
	//
	// ops-cn-em93wcq0s001
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
}

func (s DeleteWorkspaceResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s DeleteWorkspaceResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DeleteWorkspaceResponseBodyResult) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteWorkspaceResponseBodyResult) SetInstanceId(v string) *DeleteWorkspaceResponseBodyResult {
	s.InstanceId = &v
	return s
}

func (s *DeleteWorkspaceResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
