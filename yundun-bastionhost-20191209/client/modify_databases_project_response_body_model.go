// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDatabasesProjectResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyDatabasesProjectResponseBody
	GetRequestId() *string
	SetResults(v []*ModifyDatabasesProjectResponseBodyResults) *ModifyDatabasesProjectResponseBody
	GetResults() []*ModifyDatabasesProjectResponseBodyResults
	SetSuccessCount(v int64) *ModifyDatabasesProjectResponseBody
	GetSuccessCount() *int64
	SetTotalCount(v int64) *ModifyDatabasesProjectResponseBody
	GetTotalCount() *int64
}

type ModifyDatabasesProjectResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// EC9BF0F4-8983-491A-BC8C-1B4DD94976DE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The results of the operation.
	Results []*ModifyDatabasesProjectResponseBodyResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
	// The number of successful operations.
	//
	// example:
	//
	// 10
	SuccessCount *int64 `json:"SuccessCount,omitempty" xml:"SuccessCount,omitempty"`
	// The total number of records.
	//
	// example:
	//
	// 15
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ModifyDatabasesProjectResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyDatabasesProjectResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDatabasesProjectResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyDatabasesProjectResponseBody) GetResults() []*ModifyDatabasesProjectResponseBodyResults {
	return s.Results
}

func (s *ModifyDatabasesProjectResponseBody) GetSuccessCount() *int64 {
	return s.SuccessCount
}

func (s *ModifyDatabasesProjectResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ModifyDatabasesProjectResponseBody) SetRequestId(v string) *ModifyDatabasesProjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDatabasesProjectResponseBody) SetResults(v []*ModifyDatabasesProjectResponseBodyResults) *ModifyDatabasesProjectResponseBody {
	s.Results = v
	return s
}

func (s *ModifyDatabasesProjectResponseBody) SetSuccessCount(v int64) *ModifyDatabasesProjectResponseBody {
	s.SuccessCount = &v
	return s
}

func (s *ModifyDatabasesProjectResponseBody) SetTotalCount(v int64) *ModifyDatabasesProjectResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ModifyDatabasesProjectResponseBody) Validate() error {
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

type ModifyDatabasesProjectResponseBodyResults struct {
	// The database instance ID.
	//
	// example:
	//
	// 53
	AssetId *string `json:"AssetId,omitempty" xml:"AssetId,omitempty"`
	// The error code. An error code of **200*	- indicates that the operation is successful. Other error codes indicate that the operation failed.
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

func (s ModifyDatabasesProjectResponseBodyResults) String() string {
	return dara.Prettify(s)
}

func (s ModifyDatabasesProjectResponseBodyResults) GoString() string {
	return s.String()
}

func (s *ModifyDatabasesProjectResponseBodyResults) GetAssetId() *string {
	return s.AssetId
}

func (s *ModifyDatabasesProjectResponseBodyResults) GetCode() *string {
	return s.Code
}

func (s *ModifyDatabasesProjectResponseBodyResults) GetMessage() *string {
	return s.Message
}

func (s *ModifyDatabasesProjectResponseBodyResults) SetAssetId(v string) *ModifyDatabasesProjectResponseBodyResults {
	s.AssetId = &v
	return s
}

func (s *ModifyDatabasesProjectResponseBodyResults) SetCode(v string) *ModifyDatabasesProjectResponseBodyResults {
	s.Code = &v
	return s
}

func (s *ModifyDatabasesProjectResponseBodyResults) SetMessage(v string) *ModifyDatabasesProjectResponseBodyResults {
	s.Message = &v
	return s
}

func (s *ModifyDatabasesProjectResponseBodyResults) Validate() error {
	return dara.Validate(s)
}
