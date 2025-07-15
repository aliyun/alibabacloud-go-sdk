// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveTerminalsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RemoveTerminalsResponseBody
	GetRequestId() *string
	SetTerminals(v []*RemoveTerminalsResponseBodyTerminals) *RemoveTerminalsResponseBody
	GetTerminals() []*RemoveTerminalsResponseBodyTerminals
}

type RemoveTerminalsResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 16A96B9A-F203-4EC5-8E43-CB92E68F4AF8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the users.
	Terminals []*RemoveTerminalsResponseBodyTerminals `json:"Terminals,omitempty" xml:"Terminals,omitempty" type:"Repeated"`
}

func (s RemoveTerminalsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RemoveTerminalsResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveTerminalsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RemoveTerminalsResponseBody) GetTerminals() []*RemoveTerminalsResponseBodyTerminals {
	return s.Terminals
}

func (s *RemoveTerminalsResponseBody) SetRequestId(v string) *RemoveTerminalsResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveTerminalsResponseBody) SetTerminals(v []*RemoveTerminalsResponseBodyTerminals) *RemoveTerminalsResponseBody {
	s.Terminals = v
	return s
}

func (s *RemoveTerminalsResponseBody) Validate() error {
	return dara.Validate(s)
}

type RemoveTerminalsResponseBodyTerminals struct {
	// The returned status code. A value of 0 indicates that the request is successful. If the request fails, an error message is returned.
	//
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The ID of the user.
	//
	// example:
	//
	// 1811****
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The result of removing the specified users from the channel. Valid values:
	//
	// 	- Success
	//
	// 	- Failed
	//
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s RemoveTerminalsResponseBodyTerminals) String() string {
	return dara.Prettify(s)
}

func (s RemoveTerminalsResponseBodyTerminals) GoString() string {
	return s.String()
}

func (s *RemoveTerminalsResponseBodyTerminals) GetCode() *int32 {
	return s.Code
}

func (s *RemoveTerminalsResponseBodyTerminals) GetId() *string {
	return s.Id
}

func (s *RemoveTerminalsResponseBodyTerminals) GetMessage() *string {
	return s.Message
}

func (s *RemoveTerminalsResponseBodyTerminals) SetCode(v int32) *RemoveTerminalsResponseBodyTerminals {
	s.Code = &v
	return s
}

func (s *RemoveTerminalsResponseBodyTerminals) SetId(v string) *RemoveTerminalsResponseBodyTerminals {
	s.Id = &v
	return s
}

func (s *RemoveTerminalsResponseBodyTerminals) SetMessage(v string) *RemoveTerminalsResponseBodyTerminals {
	s.Message = &v
	return s
}

func (s *RemoveTerminalsResponseBodyTerminals) Validate() error {
	return dara.Validate(s)
}
