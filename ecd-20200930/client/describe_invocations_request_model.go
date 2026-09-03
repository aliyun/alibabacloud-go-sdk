// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInvocationsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCommandType(v string) *DescribeInvocationsRequest
	GetCommandType() *string
	SetContentEncoding(v string) *DescribeInvocationsRequest
	GetContentEncoding() *string
	SetDesktopId(v string) *DescribeInvocationsRequest
	GetDesktopId() *string
	SetDesktopIds(v []*string) *DescribeInvocationsRequest
	GetDesktopIds() []*string
	SetEndUserId(v string) *DescribeInvocationsRequest
	GetEndUserId() *string
	SetIncludeInvokeDesktops(v bool) *DescribeInvocationsRequest
	GetIncludeInvokeDesktops() *bool
	SetIncludeOutput(v bool) *DescribeInvocationsRequest
	GetIncludeOutput() *bool
	SetInvokeId(v string) *DescribeInvocationsRequest
	GetInvokeId() *string
	SetInvokeStatus(v string) *DescribeInvocationsRequest
	GetInvokeStatus() *string
	SetMaxResults(v int32) *DescribeInvocationsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeInvocationsRequest
	GetNextToken() *string
	SetRegionId(v string) *DescribeInvocationsRequest
	GetRegionId() *string
}

type DescribeInvocationsRequest struct {
	// The command type of the O&M script.
	//
	// example:
	//
	// RunPowerShellScript
	CommandType *string `json:"CommandType,omitempty" xml:"CommandType,omitempty"`
	// The encoding method of the returned data.
	//
	// example:
	//
	// PlainText
	ContentEncoding *string `json:"ContentEncoding,omitempty" xml:"ContentEncoding,omitempty"`
	// The cloud desktop ID. If you specify a cloud desktop, all script execution records of the cloud desktop are queried.
	//
	// example:
	//
	// ecd-7w78ozhjcwa3u****
	DesktopId *string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty"`
	// The list of cloud desktop IDs.
	//
	// > The `DesktopId` parameter will be deprecated. Use this parameter to pass the list of cloud desktop IDs.
	DesktopIds []*string `json:"DesktopIds,omitempty" xml:"DesktopIds,omitempty" type:"Repeated"`
	// The user ID.
	//
	// example:
	//
	// alice
	EndUserId *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	// Specifies whether to return the execution results of all cloud desktops when the remote command is executed on multiple cloud desktops.
	IncludeInvokeDesktops *bool `json:"IncludeInvokeDesktops,omitempty" xml:"IncludeInvokeDesktops,omitempty"`
	// Specifies whether to return the output information of the script execution in the results.
	//
	// example:
	//
	// false
	IncludeOutput *bool `json:"IncludeOutput,omitempty" xml:"IncludeOutput,omitempty"`
	// The execution ID of the script process. Obtained from the response of [RunCommand](~~RunCommand~~).
	//
	// example:
	//
	// t-hz0jdfwd9f****
	InvokeId *string `json:"InvokeId,omitempty" xml:"InvokeId,omitempty"`
	// The overall execution status of the script. The overall execution status depends on the combined execution status of one or more cloud desktops in the execution.
	//
	// example:
	//
	// Finished
	InvokeStatus *string `json:"InvokeStatus,omitempty" xml:"InvokeStatus,omitempty"`
	// The number of entries per page for a paged query.
	//
	// - Maximum value: 50.
	//
	// - Default value: 10.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token. Set this parameter to the NextToken value returned in the previous API call.
	//
	// example:
	//
	// AAAAAV3MpHK1AP0pfERHZN5pu6nmB7qrRFJ8vmttjxPL****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The region ID. You can call [DescribeRegions](~~DescribeRegions~~) to query the regions supported by Elastic Desktop Service.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeInvocationsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeInvocationsRequest) GoString() string {
	return s.String()
}

func (s *DescribeInvocationsRequest) GetCommandType() *string {
	return s.CommandType
}

func (s *DescribeInvocationsRequest) GetContentEncoding() *string {
	return s.ContentEncoding
}

func (s *DescribeInvocationsRequest) GetDesktopId() *string {
	return s.DesktopId
}

func (s *DescribeInvocationsRequest) GetDesktopIds() []*string {
	return s.DesktopIds
}

func (s *DescribeInvocationsRequest) GetEndUserId() *string {
	return s.EndUserId
}

func (s *DescribeInvocationsRequest) GetIncludeInvokeDesktops() *bool {
	return s.IncludeInvokeDesktops
}

func (s *DescribeInvocationsRequest) GetIncludeOutput() *bool {
	return s.IncludeOutput
}

func (s *DescribeInvocationsRequest) GetInvokeId() *string {
	return s.InvokeId
}

func (s *DescribeInvocationsRequest) GetInvokeStatus() *string {
	return s.InvokeStatus
}

func (s *DescribeInvocationsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeInvocationsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeInvocationsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeInvocationsRequest) SetCommandType(v string) *DescribeInvocationsRequest {
	s.CommandType = &v
	return s
}

func (s *DescribeInvocationsRequest) SetContentEncoding(v string) *DescribeInvocationsRequest {
	s.ContentEncoding = &v
	return s
}

func (s *DescribeInvocationsRequest) SetDesktopId(v string) *DescribeInvocationsRequest {
	s.DesktopId = &v
	return s
}

func (s *DescribeInvocationsRequest) SetDesktopIds(v []*string) *DescribeInvocationsRequest {
	s.DesktopIds = v
	return s
}

func (s *DescribeInvocationsRequest) SetEndUserId(v string) *DescribeInvocationsRequest {
	s.EndUserId = &v
	return s
}

func (s *DescribeInvocationsRequest) SetIncludeInvokeDesktops(v bool) *DescribeInvocationsRequest {
	s.IncludeInvokeDesktops = &v
	return s
}

func (s *DescribeInvocationsRequest) SetIncludeOutput(v bool) *DescribeInvocationsRequest {
	s.IncludeOutput = &v
	return s
}

func (s *DescribeInvocationsRequest) SetInvokeId(v string) *DescribeInvocationsRequest {
	s.InvokeId = &v
	return s
}

func (s *DescribeInvocationsRequest) SetInvokeStatus(v string) *DescribeInvocationsRequest {
	s.InvokeStatus = &v
	return s
}

func (s *DescribeInvocationsRequest) SetMaxResults(v int32) *DescribeInvocationsRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeInvocationsRequest) SetNextToken(v string) *DescribeInvocationsRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeInvocationsRequest) SetRegionId(v string) *DescribeInvocationsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeInvocationsRequest) Validate() error {
	return dara.Validate(s)
}
