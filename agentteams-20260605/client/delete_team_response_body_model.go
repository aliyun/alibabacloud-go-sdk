// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTeamResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteTeamResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *DeleteTeamResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DeleteTeamResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteTeamResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteTeamResponseBody
	GetSuccess() *bool
}

type DeleteTeamResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteTeamResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteTeamResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteTeamResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteTeamResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteTeamResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteTeamResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteTeamResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteTeamResponseBody) SetCode(v string) *DeleteTeamResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteTeamResponseBody) SetHttpStatusCode(v int32) *DeleteTeamResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteTeamResponseBody) SetMessage(v string) *DeleteTeamResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteTeamResponseBody) SetRequestId(v string) *DeleteTeamResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteTeamResponseBody) SetSuccess(v bool) *DeleteTeamResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteTeamResponseBody) Validate() error {
	return dara.Validate(s)
}
