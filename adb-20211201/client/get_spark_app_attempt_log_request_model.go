// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSparkAppAttemptLogRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAttemptId(v string) *GetSparkAppAttemptLogRequest
	GetAttemptId() *string
	SetLogLength(v int64) *GetSparkAppAttemptLogRequest
	GetLogLength() *int64
	SetPageNumber(v int32) *GetSparkAppAttemptLogRequest
	GetPageNumber() *int32
	SetPageSize(v string) *GetSparkAppAttemptLogRequest
	GetPageSize() *string
}

type GetSparkAppAttemptLogRequest struct {
	// The ID of the log.
	//
	// > You can call the [ListSparkAppAttempts](https://help.aliyun.com/document_detail/455887.html) operation to query the information about the retry attempts of a Spark application, including the retry log IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// s202207151211hz****-0001
	AttemptId *string `json:"AttemptId,omitempty" xml:"AttemptId,omitempty"`
	// The number of log entries to return. Valid values: 1 to 500. Default value: 300.
	//
	// example:
	//
	// 20
	LogLength *int64 `json:"LogLength,omitempty" xml:"LogLength,omitempty"`
	// The log offset.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 500
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s GetSparkAppAttemptLogRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSparkAppAttemptLogRequest) GoString() string {
	return s.String()
}

func (s *GetSparkAppAttemptLogRequest) GetAttemptId() *string {
	return s.AttemptId
}

func (s *GetSparkAppAttemptLogRequest) GetLogLength() *int64 {
	return s.LogLength
}

func (s *GetSparkAppAttemptLogRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *GetSparkAppAttemptLogRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *GetSparkAppAttemptLogRequest) SetAttemptId(v string) *GetSparkAppAttemptLogRequest {
	s.AttemptId = &v
	return s
}

func (s *GetSparkAppAttemptLogRequest) SetLogLength(v int64) *GetSparkAppAttemptLogRequest {
	s.LogLength = &v
	return s
}

func (s *GetSparkAppAttemptLogRequest) SetPageNumber(v int32) *GetSparkAppAttemptLogRequest {
	s.PageNumber = &v
	return s
}

func (s *GetSparkAppAttemptLogRequest) SetPageSize(v string) *GetSparkAppAttemptLogRequest {
	s.PageSize = &v
	return s
}

func (s *GetSparkAppAttemptLogRequest) Validate() error {
	return dara.Validate(s)
}
