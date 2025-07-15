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
	// The command type.
	//
	// Valid values:
	//
	// 	- RunPowerShellScript: the PowerShell command.
	//
	// 	- RunBatScript: the Bat command.
	//
	// example:
	//
	// RunPowerShellScript
	CommandType *string `json:"CommandType,omitempty" xml:"CommandType,omitempty"`
	// The encoding method of the command content and outputs.
	//
	// Valid values:
	//
	// 	- Base64 (default): returns the Base64-encoded command content and command outputs.
	//
	// 	- PlainText: returns the original command content and outputs in plain text.
	//
	// example:
	//
	// PlainText
	ContentEncoding *string `json:"ContentEncoding,omitempty" xml:"ContentEncoding,omitempty"`
	// The cloud computer ID. If you specify a cloud computer, all command execution records of the cloud computer are queried.
	//
	// example:
	//
	// ecd-7w78ozhjcwa3u****
	DesktopId *string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty"`
	// The cloud computer IDs.
	//
	// >  The `DesktopId` parameter will be deprecated. We recommend using the DesktopIds parameter to specify cloud computer IDs instead.
	DesktopIds []*string `json:"DesktopIds,omitempty" xml:"DesktopIds,omitempty" type:"Repeated"`
	// The user ID.
	//
	// example:
	//
	// test1
	EndUserId *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	// Specifies whether to return the execution results of the remote command on all cloud computers when executed across multiple cloud computers.
	//
	// Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// false
	IncludeInvokeDesktops *bool `json:"IncludeInvokeDesktops,omitempty" xml:"IncludeInvokeDesktops,omitempty"`
	// Specifies whether to return command outputs in the response.
	//
	// Valid values:
	//
	// 	- true
	//
	// 	- false (default)
	//
	// example:
	//
	// false
	IncludeOutput *bool `json:"IncludeOutput,omitempty" xml:"IncludeOutput,omitempty"`
	// The execution ID of the command. You can obtain the value by calling the [RunCommand](~~RunCommand~~) operation.
	//
	// example:
	//
	// t-hz0jdfwd9f****
	InvokeId *string `json:"InvokeId,omitempty" xml:"InvokeId,omitempty"`
	// The execution status of the command. The value of this parameter is determined by the execution states of the command on all participating cloud computers.
	//
	// Valid values:
	//
	// 	- Finished: The command execution completes on all cloud computers. Alternatively, the command execution is manually stopped on some cloud computers while it completes on the others.
	//
	// 	- Stopped: The command execution stops.
	//
	// 	- Failed: The command execution failed on all cloud computers.
	//
	// 	- Running: Once there is a command execution in progress, the execution status defaults to Running.
	//
	// 	- PartialFailed: If the command execution failed on part of the cloud computers, the execution status is considered partially failed.
	//
	// example:
	//
	// Finished
	InvokeStatus *string `json:"InvokeStatus,omitempty" xml:"InvokeStatus,omitempty"`
	// The number of entries per page.
	//
	// 	- Valid values: 1 to 50.
	//
	// 	- Default value: 10.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The query token. Set the value to the NextToken value that is returned from the last call to the previous DescribeInvocations operation.
	//
	// example:
	//
	// AAAAAV3MpHK1AP0pfERHZN5pu6nmB7qrRFJ8vmttjxPL****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The region ID. You can call the [DescribeRegions](~~DescribeRegions~~) operation to query the list of regions where Elastic Desktop Service (EDS) Enterprise is available.
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
