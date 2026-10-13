// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeProtocolServiceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DescribeProtocolServiceRequest
	GetClientToken() *string
	SetDescription(v string) *DescribeProtocolServiceRequest
	GetDescription() *string
	SetFileSystemId(v string) *DescribeProtocolServiceRequest
	GetFileSystemId() *string
	SetMaxResults(v int64) *DescribeProtocolServiceRequest
	GetMaxResults() *int64
	SetNextToken(v string) *DescribeProtocolServiceRequest
	GetNextToken() *string
	SetProtocolServiceIds(v string) *DescribeProtocolServiceRequest
	GetProtocolServiceIds() *string
	SetStatus(v string) *DescribeProtocolServiceRequest
	GetStatus() *string
}

type DescribeProtocolServiceRequest struct {
	// Ensures the idempotency of the request. You can use the client to generate the token, but make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotency](https://help.aliyun.com/document_detail/25693.html).
	//
	// >If you do not specify this parameter, the system automatically uses the RequestId of the API request as the ClientToken. The RequestId may vary for each API request.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The description of the protocol service or a part of the description used for matching.
	//
	// Limits:
	//
	// - The description must be 2 to 128 characters in length.
	//
	// - The description must start with a letter and cannot start with `http://` or `https://`.
	//
	// - The description can contain digits, colons (:), underscores (_), or hyphens (-).
	//
	// example:
	//
	// Description of this protocol service
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The file system ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cpfs-099394bd928c****
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// The maximum number of results to return for each request.
	//
	// - Maximum value: 100.
	//
	// - Minimum value: 10.
	//
	// - Default value: 20.
	//
	// example:
	//
	// 20
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token used to initiate the next request if the response of the current request is truncated. You can use this token to retrieve the remaining results.
	//
	// example:
	//
	// aBcdg==
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The protocol service IDs.
	//
	// - Format: CSV
	//
	// - Limit: You can specify up to 10 protocol service IDs.
	//
	// example:
	//
	// ptc-197ed6a00f2b****,ptc-196ed6a00f2b****
	ProtocolServiceIds *string `json:"ProtocolServiceIds,omitempty" xml:"ProtocolServiceIds,omitempty"`
	// The status of the protocol service used for matching.
	//
	// Format: CSV
	//
	// Valid values:
	//
	// - Creating: The protocol service is being created.
	//
	// - Starting: The protocol service is being started.
	//
	// - Running: The protocol service is available.
	//
	// - Updating: The protocol service is being updated.
	//
	// - Deleting: The protocol service is being deleted.
	//
	// - Stopping: The protocol service is being stopped.
	//
	// - Stopped: The protocol service is stopped.
	//
	// example:
	//
	// Running,Updating
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeProtocolServiceRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeProtocolServiceRequest) GoString() string {
	return s.String()
}

func (s *DescribeProtocolServiceRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DescribeProtocolServiceRequest) GetDescription() *string {
	return s.Description
}

func (s *DescribeProtocolServiceRequest) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *DescribeProtocolServiceRequest) GetMaxResults() *int64 {
	return s.MaxResults
}

func (s *DescribeProtocolServiceRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeProtocolServiceRequest) GetProtocolServiceIds() *string {
	return s.ProtocolServiceIds
}

func (s *DescribeProtocolServiceRequest) GetStatus() *string {
	return s.Status
}

func (s *DescribeProtocolServiceRequest) SetClientToken(v string) *DescribeProtocolServiceRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribeProtocolServiceRequest) SetDescription(v string) *DescribeProtocolServiceRequest {
	s.Description = &v
	return s
}

func (s *DescribeProtocolServiceRequest) SetFileSystemId(v string) *DescribeProtocolServiceRequest {
	s.FileSystemId = &v
	return s
}

func (s *DescribeProtocolServiceRequest) SetMaxResults(v int64) *DescribeProtocolServiceRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeProtocolServiceRequest) SetNextToken(v string) *DescribeProtocolServiceRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeProtocolServiceRequest) SetProtocolServiceIds(v string) *DescribeProtocolServiceRequest {
	s.ProtocolServiceIds = &v
	return s
}

func (s *DescribeProtocolServiceRequest) SetStatus(v string) *DescribeProtocolServiceRequest {
	s.Status = &v
	return s
}

func (s *DescribeProtocolServiceRequest) Validate() error {
	return dara.Validate(s)
}
