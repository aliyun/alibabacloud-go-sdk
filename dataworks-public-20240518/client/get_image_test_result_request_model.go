// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetImageTestResultRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *GetImageTestResultRequest
	GetId() *string
	SetProcessId(v string) *GetImageTestResultRequest
	GetProcessId() *string
}

type GetImageTestResultRequest struct {
	// The image ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// Custom_image_xxxx_xxxx
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The test process ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 582d4896-d224-413b-b883-239eeebe0bc5
	ProcessId *string `json:"ProcessId,omitempty" xml:"ProcessId,omitempty"`
}

func (s GetImageTestResultRequest) String() string {
	return dara.Prettify(s)
}

func (s GetImageTestResultRequest) GoString() string {
	return s.String()
}

func (s *GetImageTestResultRequest) GetId() *string {
	return s.Id
}

func (s *GetImageTestResultRequest) GetProcessId() *string {
	return s.ProcessId
}

func (s *GetImageTestResultRequest) SetId(v string) *GetImageTestResultRequest {
	s.Id = &v
	return s
}

func (s *GetImageTestResultRequest) SetProcessId(v string) *GetImageTestResultRequest {
	s.ProcessId = &v
	return s
}

func (s *GetImageTestResultRequest) Validate() error {
	return dara.Validate(s)
}
