// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRestoreCapacityResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *RestoreCapacityResponseBody
	GetData() *bool
	SetRequestId(v string) *RestoreCapacityResponseBody
	GetRequestId() *string
}

type RestoreCapacityResponseBody struct {
	// Indicates whether the delete command was sent. Valid values:
	//
	// - true: The delete command was sent and the cleanup is in progress.
	//
	// - false: The command failed to send.
	//
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 6276D891-58D4-55B2-87B9-74D413F7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RestoreCapacityResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RestoreCapacityResponseBody) GoString() string {
	return s.String()
}

func (s *RestoreCapacityResponseBody) GetData() *bool {
	return s.Data
}

func (s *RestoreCapacityResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RestoreCapacityResponseBody) SetData(v bool) *RestoreCapacityResponseBody {
	s.Data = &v
	return s
}

func (s *RestoreCapacityResponseBody) SetRequestId(v string) *RestoreCapacityResponseBody {
	s.RequestId = &v
	return s
}

func (s *RestoreCapacityResponseBody) Validate() error {
	return dara.Validate(s)
}
