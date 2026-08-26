// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetLiveEdgeTransferRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *SetLiveEdgeTransferRequest
	GetAppName() *string
	SetDomainName(v string) *SetLiveEdgeTransferRequest
	GetDomainName() *string
	SetHttpDns(v string) *SetLiveEdgeTransferRequest
	GetHttpDns() *string
	SetOwnerId(v int64) *SetLiveEdgeTransferRequest
	GetOwnerId() *int64
	SetRegionId(v string) *SetLiveEdgeTransferRequest
	GetRegionId() *string
	SetStreamName(v string) *SetLiveEdgeTransferRequest
	GetStreamName() *string
	SetTargetDomainList(v string) *SetLiveEdgeTransferRequest
	GetTargetDomainList() *string
	SetTransferArgs(v string) *SetLiveEdgeTransferRequest
	GetTransferArgs() *string
}

type SetLiveEdgeTransferRequest struct {
	// The application name to which the live stream belongs. Regular expressions are supported for configuration with exceptions. For more information, see **AppName and StreamName Parameter Configuration Instructions*	- below. For example: liveApp****[1,2,3] indicates that the three apps liveApp****1, liveApp****2, and liveApp****3 are allowed for stream relay.
	//
	// > - This parameter only takes effect for the TargetDomainList in the request parameters.
	//
	// > - When configuring the `AppName` parameter value using regular expressions, the ^ or $ characters cannot be used, otherwise stream relay will fail.
	//
	// example:
	//
	// liveApp****
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The ingest domain name. Live stream relay is configured at the granularity of the ingest DomainName. Each domain can have only one live stream relay configuration.
	//
	// This parameter is required.
	//
	// example:
	//
	// demo.aliyundoc.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The HTTPDNS interface for obtaining the stream relay target address. The request must contain one of the `TargetDomainList` and `HttpDns` parameters, and the two are mutually exclusive.
	//
	// > If `HttpDns` is set in the request parameters, the `TargetDomainList` parameter cannot be set, and the `AppName` and `StreamName` restrictions do not take effect.
	//
	// Live stream relay has requirements for the message structure returned by the HTTPDNS interface. For more information, see **HTTPDNS Instructions*	- below.
	//
	// example:
	//
	// http://developer.aliyundoc.com
	HttpDns *string `json:"HttpDns,omitempty" xml:"HttpDns,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The stream name. Regular expressions are supported for configuration with exceptions. For more information, see **AppName and StreamName Parameter Configuration Instructions*	- below. For example: liveStream****[1,2,3] indicates that the three streams liveStream****1, liveStream****2, and liveStream****3 are allowed for stream relay.
	//
	// > - This parameter only takes effect for the TargetDomainList in the request parameters.
	//
	// > - When configuring the `StreamName` parameter value using regular expressions, the ^ or $ characters cannot be used, otherwise stream relay will fail.
	//
	// example:
	//
	// liveStream****
	StreamName *string `json:"StreamName,omitempty" xml:"StreamName,omitempty"`
	// The list of target domains specified by the user for stream relay. Multiple domains are separated by commas (,). The request must contain one of the `TargetDomainList` and `HttpDns` parameters, and the two are mutually exclusive.
	//
	// > - When `TargetDomainList` is set in the request parameters, the `AppName` and `StreamName` parameters take effect.
	//
	// > - When `TargetDomainList` is set in the request parameters, the `HttpDns` parameter cannot be set.
	//
	// example:
	//
	// learn.aliyundoc.com,guide.aliyundoc.com
	TargetDomainList *string `json:"TargetDomainList,omitempty" xml:"TargetDomainList,omitempty"`
	// Specifies whether to pass through ingest parameters. Valid values:
	//
	// - **yes**: Ingest parameters are passed through.
	//
	// - **no*	- (default): Ingest parameters are not passed through.
	//
	// example:
	//
	// yes
	TransferArgs *string `json:"TransferArgs,omitempty" xml:"TransferArgs,omitempty"`
}

func (s SetLiveEdgeTransferRequest) String() string {
	return dara.Prettify(s)
}

func (s SetLiveEdgeTransferRequest) GoString() string {
	return s.String()
}

func (s *SetLiveEdgeTransferRequest) GetAppName() *string {
	return s.AppName
}

func (s *SetLiveEdgeTransferRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *SetLiveEdgeTransferRequest) GetHttpDns() *string {
	return s.HttpDns
}

func (s *SetLiveEdgeTransferRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *SetLiveEdgeTransferRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *SetLiveEdgeTransferRequest) GetStreamName() *string {
	return s.StreamName
}

func (s *SetLiveEdgeTransferRequest) GetTargetDomainList() *string {
	return s.TargetDomainList
}

func (s *SetLiveEdgeTransferRequest) GetTransferArgs() *string {
	return s.TransferArgs
}

func (s *SetLiveEdgeTransferRequest) SetAppName(v string) *SetLiveEdgeTransferRequest {
	s.AppName = &v
	return s
}

func (s *SetLiveEdgeTransferRequest) SetDomainName(v string) *SetLiveEdgeTransferRequest {
	s.DomainName = &v
	return s
}

func (s *SetLiveEdgeTransferRequest) SetHttpDns(v string) *SetLiveEdgeTransferRequest {
	s.HttpDns = &v
	return s
}

func (s *SetLiveEdgeTransferRequest) SetOwnerId(v int64) *SetLiveEdgeTransferRequest {
	s.OwnerId = &v
	return s
}

func (s *SetLiveEdgeTransferRequest) SetRegionId(v string) *SetLiveEdgeTransferRequest {
	s.RegionId = &v
	return s
}

func (s *SetLiveEdgeTransferRequest) SetStreamName(v string) *SetLiveEdgeTransferRequest {
	s.StreamName = &v
	return s
}

func (s *SetLiveEdgeTransferRequest) SetTargetDomainList(v string) *SetLiveEdgeTransferRequest {
	s.TargetDomainList = &v
	return s
}

func (s *SetLiveEdgeTransferRequest) SetTransferArgs(v string) *SetLiveEdgeTransferRequest {
	s.TransferArgs = &v
	return s
}

func (s *SetLiveEdgeTransferRequest) Validate() error {
	return dara.Validate(s)
}
