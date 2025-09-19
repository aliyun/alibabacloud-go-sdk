// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOperationSuspEventsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessCode(v string) *OperationSuspEventsResponseBody
	GetAccessCode() *string
	SetRequestId(v string) *OperationSuspEventsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *OperationSuspEventsResponseBody
	GetSuccess() *bool
}

type OperationSuspEventsResponseBody struct {
	// Indicates whether you have access permissions. Valid values:
	//
	// 	- **pass**: yes
	//
	// 	- **no_permission**: no
	//
	// example:
	//
	// pass
	AccessCode *string `json:"AccessCode,omitempty" xml:"AccessCode,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 7E0618A9-D5EF-4220-9471-C42B5E92719F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether exceptions are handled. Valid values:
	//
	// 	- **true**: yes
	//
	// 	- **false**: no
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s OperationSuspEventsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s OperationSuspEventsResponseBody) GoString() string {
	return s.String()
}

func (s *OperationSuspEventsResponseBody) GetAccessCode() *string {
	return s.AccessCode
}

func (s *OperationSuspEventsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *OperationSuspEventsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *OperationSuspEventsResponseBody) SetAccessCode(v string) *OperationSuspEventsResponseBody {
	s.AccessCode = &v
	return s
}

func (s *OperationSuspEventsResponseBody) SetRequestId(v string) *OperationSuspEventsResponseBody {
	s.RequestId = &v
	return s
}

func (s *OperationSuspEventsResponseBody) SetSuccess(v bool) *OperationSuspEventsResponseBody {
	s.Success = &v
	return s
}

func (s *OperationSuspEventsResponseBody) Validate() error {
	return dara.Validate(s)
}
