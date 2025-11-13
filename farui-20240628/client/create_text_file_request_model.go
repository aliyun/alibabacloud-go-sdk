// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTextFileRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateTextFileRequest
	GetClientToken() *string
	SetContractId(v string) *CreateTextFileRequest
	GetContractId() *string
	SetCreateTime(v string) *CreateTextFileRequest
	GetCreateTime() *string
	SetTextFileName(v string) *CreateTextFileRequest
	GetTextFileName() *string
	SetTextFileUrl(v string) *CreateTextFileRequest
	GetTextFileUrl() *string
}

type CreateTextFileRequest struct {
	// example:
	//
	// e9a93201-7e96-4dc1-9678-2832fc132d08
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	ContractId  *string `json:"ContractId,omitempty" xml:"ContractId,omitempty"`
	// example:
	//
	// 1714476549
	CreateTime   *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	TextFileName *string `json:"TextFileName,omitempty" xml:"TextFileName,omitempty"`
	TextFileUrl  *string `json:"TextFileUrl,omitempty" xml:"TextFileUrl,omitempty"`
}

func (s CreateTextFileRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateTextFileRequest) GoString() string {
	return s.String()
}

func (s *CreateTextFileRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateTextFileRequest) GetContractId() *string {
	return s.ContractId
}

func (s *CreateTextFileRequest) GetCreateTime() *string {
	return s.CreateTime
}

func (s *CreateTextFileRequest) GetTextFileName() *string {
	return s.TextFileName
}

func (s *CreateTextFileRequest) GetTextFileUrl() *string {
	return s.TextFileUrl
}

func (s *CreateTextFileRequest) SetClientToken(v string) *CreateTextFileRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateTextFileRequest) SetContractId(v string) *CreateTextFileRequest {
	s.ContractId = &v
	return s
}

func (s *CreateTextFileRequest) SetCreateTime(v string) *CreateTextFileRequest {
	s.CreateTime = &v
	return s
}

func (s *CreateTextFileRequest) SetTextFileName(v string) *CreateTextFileRequest {
	s.TextFileName = &v
	return s
}

func (s *CreateTextFileRequest) SetTextFileUrl(v string) *CreateTextFileRequest {
	s.TextFileUrl = &v
	return s
}

func (s *CreateTextFileRequest) Validate() error {
	return dara.Validate(s)
}
