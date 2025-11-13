// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
	"io"
)

type iCreateTextFileAdvanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateTextFileAdvanceRequest
	GetClientToken() *string
	SetContractId(v string) *CreateTextFileAdvanceRequest
	GetContractId() *string
	SetCreateTime(v string) *CreateTextFileAdvanceRequest
	GetCreateTime() *string
	SetTextFileName(v string) *CreateTextFileAdvanceRequest
	GetTextFileName() *string
	SetTextFileUrlObject(v io.Reader) *CreateTextFileAdvanceRequest
	GetTextFileUrlObject() io.Reader
}

type CreateTextFileAdvanceRequest struct {
	// example:
	//
	// e9a93201-7e96-4dc1-9678-2832fc132d08
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	ContractId  *string `json:"ContractId,omitempty" xml:"ContractId,omitempty"`
	// example:
	//
	// 1714476549
	CreateTime        *string   `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	TextFileName      *string   `json:"TextFileName,omitempty" xml:"TextFileName,omitempty"`
	TextFileUrlObject io.Reader `json:"TextFileUrl,omitempty" xml:"TextFileUrl,omitempty"`
}

func (s CreateTextFileAdvanceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateTextFileAdvanceRequest) GoString() string {
	return s.String()
}

func (s *CreateTextFileAdvanceRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateTextFileAdvanceRequest) GetContractId() *string {
	return s.ContractId
}

func (s *CreateTextFileAdvanceRequest) GetCreateTime() *string {
	return s.CreateTime
}

func (s *CreateTextFileAdvanceRequest) GetTextFileName() *string {
	return s.TextFileName
}

func (s *CreateTextFileAdvanceRequest) GetTextFileUrlObject() io.Reader {
	return s.TextFileUrlObject
}

func (s *CreateTextFileAdvanceRequest) SetClientToken(v string) *CreateTextFileAdvanceRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateTextFileAdvanceRequest) SetContractId(v string) *CreateTextFileAdvanceRequest {
	s.ContractId = &v
	return s
}

func (s *CreateTextFileAdvanceRequest) SetCreateTime(v string) *CreateTextFileAdvanceRequest {
	s.CreateTime = &v
	return s
}

func (s *CreateTextFileAdvanceRequest) SetTextFileName(v string) *CreateTextFileAdvanceRequest {
	s.TextFileName = &v
	return s
}

func (s *CreateTextFileAdvanceRequest) SetTextFileUrlObject(v io.Reader) *CreateTextFileAdvanceRequest {
	s.TextFileUrlObject = v
	return s
}

func (s *CreateTextFileAdvanceRequest) Validate() error {
	return dara.Validate(s)
}
