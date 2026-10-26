// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyHostsProjectResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyHostsProjectResponseBody
	GetRequestId() *string
	SetResults(v []*ModifyHostsProjectResponseBodyResults) *ModifyHostsProjectResponseBody
	GetResults() []*ModifyHostsProjectResponseBodyResults
	SetSuccessCount(v int64) *ModifyHostsProjectResponseBody
	GetSuccessCount() *int64
	SetTotalCount(v int64) *ModifyHostsProjectResponseBody
	GetTotalCount() *int64
}

type ModifyHostsProjectResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 81500666-d7f5-4143-8329-0223cc738105
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The call results of the operation.
	Results []*ModifyHostsProjectResponseBodyResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
	// The number of hosts that were successfully moved.
	//
	// example:
	//
	// 10
	SuccessCount *int64 `json:"SuccessCount,omitempty" xml:"SuccessCount,omitempty"`
	// The total number of records.
	//
	// example:
	//
	// 10
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ModifyHostsProjectResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyHostsProjectResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyHostsProjectResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyHostsProjectResponseBody) GetResults() []*ModifyHostsProjectResponseBodyResults {
	return s.Results
}

func (s *ModifyHostsProjectResponseBody) GetSuccessCount() *int64 {
	return s.SuccessCount
}

func (s *ModifyHostsProjectResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ModifyHostsProjectResponseBody) SetRequestId(v string) *ModifyHostsProjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyHostsProjectResponseBody) SetResults(v []*ModifyHostsProjectResponseBodyResults) *ModifyHostsProjectResponseBody {
	s.Results = v
	return s
}

func (s *ModifyHostsProjectResponseBody) SetSuccessCount(v int64) *ModifyHostsProjectResponseBody {
	s.SuccessCount = &v
	return s
}

func (s *ModifyHostsProjectResponseBody) SetTotalCount(v int64) *ModifyHostsProjectResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ModifyHostsProjectResponseBody) Validate() error {
	if s.Results != nil {
		for _, item := range s.Results {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ModifyHostsProjectResponseBodyResults struct {
	// The host ID.
	//
	// example:
	//
	// 16
	AssetId *string `json:"AssetId,omitempty" xml:"AssetId,omitempty"`
	// The error code. A value of **200*	- indicates that the operation was successful. Other values indicate that the operation failed.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The result message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s ModifyHostsProjectResponseBodyResults) String() string {
	return dara.Prettify(s)
}

func (s ModifyHostsProjectResponseBodyResults) GoString() string {
	return s.String()
}

func (s *ModifyHostsProjectResponseBodyResults) GetAssetId() *string {
	return s.AssetId
}

func (s *ModifyHostsProjectResponseBodyResults) GetCode() *string {
	return s.Code
}

func (s *ModifyHostsProjectResponseBodyResults) GetMessage() *string {
	return s.Message
}

func (s *ModifyHostsProjectResponseBodyResults) SetAssetId(v string) *ModifyHostsProjectResponseBodyResults {
	s.AssetId = &v
	return s
}

func (s *ModifyHostsProjectResponseBodyResults) SetCode(v string) *ModifyHostsProjectResponseBodyResults {
	s.Code = &v
	return s
}

func (s *ModifyHostsProjectResponseBodyResults) SetMessage(v string) *ModifyHostsProjectResponseBodyResults {
	s.Message = &v
	return s
}

func (s *ModifyHostsProjectResponseBodyResults) Validate() error {
	return dara.Validate(s)
}
