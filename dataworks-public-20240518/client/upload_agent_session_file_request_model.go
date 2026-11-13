// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadAgentSessionFileRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *UploadAgentSessionFileRequest
	GetId() *string
	SetJsonrpc(v string) *UploadAgentSessionFileRequest
	GetJsonrpc() *string
	SetParams(v *UploadAgentSessionFileRequestParams) *UploadAgentSessionFileRequest
	GetParams() *UploadAgentSessionFileRequestParams
}

type UploadAgentSessionFileRequest struct {
	// example:
	//
	// attachment-rpc-001
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 2.0
	Jsonrpc *string `json:"Jsonrpc,omitempty" xml:"Jsonrpc,omitempty"`
	// This parameter is required.
	Params *UploadAgentSessionFileRequestParams `json:"Params,omitempty" xml:"Params,omitempty" type:"Struct"`
}

func (s UploadAgentSessionFileRequest) String() string {
	return dara.Prettify(s)
}

func (s UploadAgentSessionFileRequest) GoString() string {
	return s.String()
}

func (s *UploadAgentSessionFileRequest) GetId() *string {
	return s.Id
}

func (s *UploadAgentSessionFileRequest) GetJsonrpc() *string {
	return s.Jsonrpc
}

func (s *UploadAgentSessionFileRequest) GetParams() *UploadAgentSessionFileRequestParams {
	return s.Params
}

func (s *UploadAgentSessionFileRequest) SetId(v string) *UploadAgentSessionFileRequest {
	s.Id = &v
	return s
}

func (s *UploadAgentSessionFileRequest) SetJsonrpc(v string) *UploadAgentSessionFileRequest {
	s.Jsonrpc = &v
	return s
}

func (s *UploadAgentSessionFileRequest) SetParams(v *UploadAgentSessionFileRequestParams) *UploadAgentSessionFileRequest {
	s.Params = v
	return s
}

func (s *UploadAgentSessionFileRequest) Validate() error {
	if s.Params != nil {
		if err := s.Params.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UploadAgentSessionFileRequestParams struct {
	// This parameter is required.
	//
	// example:
	//
	// b1946ac92492d2347c6235b4d2611184
	ContentMd5 *string `json:"ContentMd5,omitempty" xml:"ContentMd5,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// sample.csv
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// session
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 6
	Size *int64 `json:"Size,omitempty" xml:"Size,omitempty"`
	// example:
	//
	// bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb
	UploadId *string `json:"UploadId,omitempty" xml:"UploadId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// prepare
	UploadMode *string `json:"UploadMode,omitempty" xml:"UploadMode,omitempty"`
}

func (s UploadAgentSessionFileRequestParams) String() string {
	return dara.Prettify(s)
}

func (s UploadAgentSessionFileRequestParams) GoString() string {
	return s.String()
}

func (s *UploadAgentSessionFileRequestParams) GetContentMd5() *string {
	return s.ContentMd5
}

func (s *UploadAgentSessionFileRequestParams) GetName() *string {
	return s.Name
}

func (s *UploadAgentSessionFileRequestParams) GetSessionId() *string {
	return s.SessionId
}

func (s *UploadAgentSessionFileRequestParams) GetSize() *int64 {
	return s.Size
}

func (s *UploadAgentSessionFileRequestParams) GetUploadId() *string {
	return s.UploadId
}

func (s *UploadAgentSessionFileRequestParams) GetUploadMode() *string {
	return s.UploadMode
}

func (s *UploadAgentSessionFileRequestParams) SetContentMd5(v string) *UploadAgentSessionFileRequestParams {
	s.ContentMd5 = &v
	return s
}

func (s *UploadAgentSessionFileRequestParams) SetName(v string) *UploadAgentSessionFileRequestParams {
	s.Name = &v
	return s
}

func (s *UploadAgentSessionFileRequestParams) SetSessionId(v string) *UploadAgentSessionFileRequestParams {
	s.SessionId = &v
	return s
}

func (s *UploadAgentSessionFileRequestParams) SetSize(v int64) *UploadAgentSessionFileRequestParams {
	s.Size = &v
	return s
}

func (s *UploadAgentSessionFileRequestParams) SetUploadId(v string) *UploadAgentSessionFileRequestParams {
	s.UploadId = &v
	return s
}

func (s *UploadAgentSessionFileRequestParams) SetUploadMode(v string) *UploadAgentSessionFileRequestParams {
	s.UploadMode = &v
	return s
}

func (s *UploadAgentSessionFileRequestParams) Validate() error {
	return dara.Validate(s)
}
