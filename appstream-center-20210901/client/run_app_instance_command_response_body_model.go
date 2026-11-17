// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunAppInstanceCommandResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCommandResultModel(v *RunAppInstanceCommandResponseBodyCommandResultModel) *RunAppInstanceCommandResponseBody
	GetCommandResultModel() *RunAppInstanceCommandResponseBodyCommandResultModel
	SetRequestId(v string) *RunAppInstanceCommandResponseBody
	GetRequestId() *string
}

type RunAppInstanceCommandResponseBody struct {
	// The command issuance result.
	CommandResultModel *RunAppInstanceCommandResponseBodyCommandResultModel `json:"CommandResultModel,omitempty" xml:"CommandResultModel,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RunAppInstanceCommandResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RunAppInstanceCommandResponseBody) GoString() string {
	return s.String()
}

func (s *RunAppInstanceCommandResponseBody) GetCommandResultModel() *RunAppInstanceCommandResponseBodyCommandResultModel {
	return s.CommandResultModel
}

func (s *RunAppInstanceCommandResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RunAppInstanceCommandResponseBody) SetCommandResultModel(v *RunAppInstanceCommandResponseBodyCommandResultModel) *RunAppInstanceCommandResponseBody {
	s.CommandResultModel = v
	return s
}

func (s *RunAppInstanceCommandResponseBody) SetRequestId(v string) *RunAppInstanceCommandResponseBody {
	s.RequestId = &v
	return s
}

func (s *RunAppInstanceCommandResponseBody) Validate() error {
	if s.CommandResultModel != nil {
		if err := s.CommandResultModel.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type RunAppInstanceCommandResponseBodyCommandResultModel struct {
	// The command execution ID. When commands are issued in batches to multiple delivery groups, only one execution ID is returned. You can use this ID to call the ListCommandInvocationResults operation to query the command execution results on each session.
	//
	// example:
	//
	// t-hz0jdfwd9f****
	InvocationId *string `json:"InvocationId,omitempty" xml:"InvocationId,omitempty"`
}

func (s RunAppInstanceCommandResponseBodyCommandResultModel) String() string {
	return dara.Prettify(s)
}

func (s RunAppInstanceCommandResponseBodyCommandResultModel) GoString() string {
	return s.String()
}

func (s *RunAppInstanceCommandResponseBodyCommandResultModel) GetInvocationId() *string {
	return s.InvocationId
}

func (s *RunAppInstanceCommandResponseBodyCommandResultModel) SetInvocationId(v string) *RunAppInstanceCommandResponseBodyCommandResultModel {
	s.InvocationId = &v
	return s
}

func (s *RunAppInstanceCommandResponseBodyCommandResultModel) Validate() error {
	return dara.Validate(s)
}
