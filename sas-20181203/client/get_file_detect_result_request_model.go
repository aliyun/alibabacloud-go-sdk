// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFileDetectResultRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHashKeyList(v []*string) *GetFileDetectResultRequest
	GetHashKeyList() []*string
	SetSourceIp(v string) *GetFileDetectResultRequest
	GetSourceIp() *string
	SetType(v int32) *GetFileDetectResultRequest
	GetType() *int32
}

type GetFileDetectResultRequest struct {
	// The identifiers of files. Only MD5 hash values are supported.
	//
	// This parameter is required.
	HashKeyList []*string `json:"HashKeyList,omitempty" xml:"HashKeyList,omitempty" type:"Repeated"`
	// The source IP address of the request.
	//
	// example:
	//
	// 183.46.XX.XX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	// The type of the file. Valid values:
	//
	// 	- **0**: unknown file
	//
	// 	- **1**: binary file
	//
	// 	- **2**: webshell file
	//
	// 	- **4**: script file
	//
	// > If you do not know the type of the file, set this parameter to 0.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0
	Type *int32 `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetFileDetectResultRequest) String() string {
	return dara.Prettify(s)
}

func (s GetFileDetectResultRequest) GoString() string {
	return s.String()
}

func (s *GetFileDetectResultRequest) GetHashKeyList() []*string {
	return s.HashKeyList
}

func (s *GetFileDetectResultRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *GetFileDetectResultRequest) GetType() *int32 {
	return s.Type
}

func (s *GetFileDetectResultRequest) SetHashKeyList(v []*string) *GetFileDetectResultRequest {
	s.HashKeyList = v
	return s
}

func (s *GetFileDetectResultRequest) SetSourceIp(v string) *GetFileDetectResultRequest {
	s.SourceIp = &v
	return s
}

func (s *GetFileDetectResultRequest) SetType(v int32) *GetFileDetectResultRequest {
	s.Type = &v
	return s
}

func (s *GetFileDetectResultRequest) Validate() error {
	return dara.Validate(s)
}
