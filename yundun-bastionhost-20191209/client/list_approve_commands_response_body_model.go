// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApproveCommandsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApproveCommands(v []*ListApproveCommandsResponseBodyApproveCommands) *ListApproveCommandsResponseBody
	GetApproveCommands() []*ListApproveCommandsResponseBodyApproveCommands
	SetRequestId(v string) *ListApproveCommandsResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListApproveCommandsResponseBody
	GetTotalCount() *int64
}

type ListApproveCommandsResponseBody struct {
	// The commands to be reviewed.
	ApproveCommands []*ListApproveCommandsResponseBodyApproveCommands `json:"ApproveCommands,omitempty" xml:"ApproveCommands,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// E3EF7711-766D-5888-997B-EFBA76809229
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of commands to be reviewed.
	//
	// example:
	//
	// 15
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListApproveCommandsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListApproveCommandsResponseBody) GoString() string {
	return s.String()
}

func (s *ListApproveCommandsResponseBody) GetApproveCommands() []*ListApproveCommandsResponseBodyApproveCommands {
	return s.ApproveCommands
}

func (s *ListApproveCommandsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListApproveCommandsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListApproveCommandsResponseBody) SetApproveCommands(v []*ListApproveCommandsResponseBodyApproveCommands) *ListApproveCommandsResponseBody {
	s.ApproveCommands = v
	return s
}

func (s *ListApproveCommandsResponseBody) SetRequestId(v string) *ListApproveCommandsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListApproveCommandsResponseBody) SetTotalCount(v int64) *ListApproveCommandsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListApproveCommandsResponseBody) Validate() error {
	if s.ApproveCommands != nil {
		for _, item := range s.ApproveCommands {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListApproveCommandsResponseBodyApproveCommands struct {
	// The ID of the command to be reviewed.
	//
	// example:
	//
	// 1
	ApproveCommandId *string `json:"ApproveCommandId,omitempty" xml:"ApproveCommandId,omitempty"`
	// The username of the asset account that is used for O\\&M.
	//
	// example:
	//
	// root
	AssetAccountName *string `json:"AssetAccountName,omitempty" xml:"AssetAccountName,omitempty"`
	// The IP address of the asset for O\\&M.
	//
	// example:
	//
	// 10.167.XX.XX
	AssetIp *string `json:"AssetIp,omitempty" xml:"AssetIp,omitempty"`
	// The name of the asset.
	//
	// example:
	//
	// poros-test
	AssetName *string `json:"AssetName,omitempty" xml:"AssetName,omitempty"`
	// The source IP address from which the application is submitted.
	//
	// example:
	//
	// 172.18.XX.XX
	ClientIp *string `json:"ClientIp,omitempty" xml:"ClientIp,omitempty"`
	// The Bastionhost user who submitted the O\\&M application.
	//
	// example:
	//
	// test
	ClientUser *string `json:"ClientUser,omitempty" xml:"ClientUser,omitempty"`
	// The command to be reviewed.
	//
	// example:
	//
	// /bin/bash
	Command *string `json:"Command,omitempty" xml:"Command,omitempty"`
	// The time when the O\\&M application was submitted. The value is a UNIX timestamp. Unit: seconds.
	//
	// example:
	//
	// 1679393152
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The O\\&M protocol.
	//
	// example:
	//
	// SSH
	ProtocolName *string `json:"ProtocolName,omitempty" xml:"ProtocolName,omitempty"`
	// The ID of the O\\&M session that triggered the review.
	//
	// example:
	//
	// 95f873ab64a76d5b0000000000004d5e
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// The status of the review. Valid values: **Wait**: The command is pending review.
	//
	// example:
	//
	// Wait
	State *string `json:"State,omitempty" xml:"State,omitempty"`
}

func (s ListApproveCommandsResponseBodyApproveCommands) String() string {
	return dara.Prettify(s)
}

func (s ListApproveCommandsResponseBodyApproveCommands) GoString() string {
	return s.String()
}

func (s *ListApproveCommandsResponseBodyApproveCommands) GetApproveCommandId() *string {
	return s.ApproveCommandId
}

func (s *ListApproveCommandsResponseBodyApproveCommands) GetAssetAccountName() *string {
	return s.AssetAccountName
}

func (s *ListApproveCommandsResponseBodyApproveCommands) GetAssetIp() *string {
	return s.AssetIp
}

func (s *ListApproveCommandsResponseBodyApproveCommands) GetAssetName() *string {
	return s.AssetName
}

func (s *ListApproveCommandsResponseBodyApproveCommands) GetClientIp() *string {
	return s.ClientIp
}

func (s *ListApproveCommandsResponseBodyApproveCommands) GetClientUser() *string {
	return s.ClientUser
}

func (s *ListApproveCommandsResponseBodyApproveCommands) GetCommand() *string {
	return s.Command
}

func (s *ListApproveCommandsResponseBodyApproveCommands) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListApproveCommandsResponseBodyApproveCommands) GetProtocolName() *string {
	return s.ProtocolName
}

func (s *ListApproveCommandsResponseBodyApproveCommands) GetSessionId() *string {
	return s.SessionId
}

func (s *ListApproveCommandsResponseBodyApproveCommands) GetState() *string {
	return s.State
}

func (s *ListApproveCommandsResponseBodyApproveCommands) SetApproveCommandId(v string) *ListApproveCommandsResponseBodyApproveCommands {
	s.ApproveCommandId = &v
	return s
}

func (s *ListApproveCommandsResponseBodyApproveCommands) SetAssetAccountName(v string) *ListApproveCommandsResponseBodyApproveCommands {
	s.AssetAccountName = &v
	return s
}

func (s *ListApproveCommandsResponseBodyApproveCommands) SetAssetIp(v string) *ListApproveCommandsResponseBodyApproveCommands {
	s.AssetIp = &v
	return s
}

func (s *ListApproveCommandsResponseBodyApproveCommands) SetAssetName(v string) *ListApproveCommandsResponseBodyApproveCommands {
	s.AssetName = &v
	return s
}

func (s *ListApproveCommandsResponseBodyApproveCommands) SetClientIp(v string) *ListApproveCommandsResponseBodyApproveCommands {
	s.ClientIp = &v
	return s
}

func (s *ListApproveCommandsResponseBodyApproveCommands) SetClientUser(v string) *ListApproveCommandsResponseBodyApproveCommands {
	s.ClientUser = &v
	return s
}

func (s *ListApproveCommandsResponseBodyApproveCommands) SetCommand(v string) *ListApproveCommandsResponseBodyApproveCommands {
	s.Command = &v
	return s
}

func (s *ListApproveCommandsResponseBodyApproveCommands) SetCreateTime(v string) *ListApproveCommandsResponseBodyApproveCommands {
	s.CreateTime = &v
	return s
}

func (s *ListApproveCommandsResponseBodyApproveCommands) SetProtocolName(v string) *ListApproveCommandsResponseBodyApproveCommands {
	s.ProtocolName = &v
	return s
}

func (s *ListApproveCommandsResponseBodyApproveCommands) SetSessionId(v string) *ListApproveCommandsResponseBodyApproveCommands {
	s.SessionId = &v
	return s
}

func (s *ListApproveCommandsResponseBodyApproveCommands) SetState(v string) *ListApproveCommandsResponseBodyApproveCommands {
	s.State = &v
	return s
}

func (s *ListApproveCommandsResponseBodyApproveCommands) Validate() error {
	return dara.Validate(s)
}
