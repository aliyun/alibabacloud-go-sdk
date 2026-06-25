// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopAndroidInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAndroidInstanceIds(v []*string) *StopAndroidInstanceRequest
	GetAndroidInstanceIds() []*string
	SetForceStop(v bool) *StopAndroidInstanceRequest
	GetForceStop() *bool
	SetSaleMode(v string) *StopAndroidInstanceRequest
	GetSaleMode() *string
}

type StopAndroidInstanceRequest struct {
	// A list of instance IDs.
	AndroidInstanceIds []*string `json:"AndroidInstanceIds,omitempty" xml:"AndroidInstanceIds,omitempty" type:"Repeated"`
	// Specifies whether to forcibly shut down the instance. If an instance cannot shut down because of a system or network exception, you can force it to shut down. This may cause data loss.
	//
	// example:
	//
	// false
	ForceStop *bool `json:"ForceStop,omitempty" xml:"ForceStop,omitempty"`
	// The sale pattern. This parameter is deprecated.
	//
	// example:
	//
	// Instance
	SaleMode *string `json:"SaleMode,omitempty" xml:"SaleMode,omitempty"`
}

func (s StopAndroidInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s StopAndroidInstanceRequest) GoString() string {
	return s.String()
}

func (s *StopAndroidInstanceRequest) GetAndroidInstanceIds() []*string {
	return s.AndroidInstanceIds
}

func (s *StopAndroidInstanceRequest) GetForceStop() *bool {
	return s.ForceStop
}

func (s *StopAndroidInstanceRequest) GetSaleMode() *string {
	return s.SaleMode
}

func (s *StopAndroidInstanceRequest) SetAndroidInstanceIds(v []*string) *StopAndroidInstanceRequest {
	s.AndroidInstanceIds = v
	return s
}

func (s *StopAndroidInstanceRequest) SetForceStop(v bool) *StopAndroidInstanceRequest {
	s.ForceStop = &v
	return s
}

func (s *StopAndroidInstanceRequest) SetSaleMode(v string) *StopAndroidInstanceRequest {
	s.SaleMode = &v
	return s
}

func (s *StopAndroidInstanceRequest) Validate() error {
	return dara.Validate(s)
}
